package pipe

import (
	"sync"
	"time"

	"github.com/minus5/go-uof-sdk"
)

type playerAPI interface {
	Player(lang uof.Lang, playerID int) (*uof.Player, []byte, error)
}

type player struct {
	api       playerAPI
	em        *expireMap
	languages []uof.Lang // suported languages
	errc      chan<- error
	out       chan<- *uof.Message
	rateLimit chan struct{}
	subProcs  *sync.WaitGroup
}

func Player(api playerAPI, languages []uof.Lang) InnerStage {
	p := &player{
		api:       api,
		languages: languages,
		em:        newExpireMap(time.Hour),
		subProcs:  &sync.WaitGroup{},
		rateLimit: make(chan struct{}, ConcurentAPICallsLimit),
	}
	return StageWithSubProcessesSync(p.loop)
}

type playerGetRequest struct {
	oddsChange  *uof.OddsChange
	requestedAt int
}

func (p *player) loop(in <-chan *uof.Message, out chan<- *uof.Message, errc chan<- error) *sync.WaitGroup {
	p.errc, p.out = errc, out

	requests := make(chan playerGetRequest, 1024)
	go func() {
		for req := range requests {
			req.oddsChange.EachPlayer(func(playerID int) {
				p.get(playerID, req.requestedAt)
			})
		}
	}()

	for m := range in {
		out <- m
		if m.Is(uof.MessageTypeOddsChange) {
			requests <- playerGetRequest{oddsChange: m.OddsChange, requestedAt: m.ReceivedAt}
		}
	}
	close(requests)
	return p.subProcs
}

func (p *player) get(playerID, requestedAt int) {
	if p.em.fresh(playerID) {
		return
	}
	p.em.insert(playerID)

	p.subProcs.Add(len(p.languages))
	for _, lang := range p.languages {
		go func(lang uof.Lang) {
			defer p.subProcs.Done()
			p.rateLimit <- struct{}{}
			defer func() { <-p.rateLimit }()

			ap, raw, err := p.api.Player(lang, playerID)
			if err != nil {
				if !uof.IsApiNotFoundErr(err) {
					p.em.remove(playerID)
				}
				p.errc <- err
				return
			}
			p.out <- uof.NewPlayerMessage(lang, ap, requestedAt, raw)
		}(lang)
	}
}
