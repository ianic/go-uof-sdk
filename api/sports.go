package api

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"

	"github.com/minus5/go-uof-sdk"
)

const (
	pathMarkets       = "/v1/descriptions/{{.Lang}}/markets.xml?include_mappings={{.IncludeMappings}}"
	pathMarketVariant = "/v1/descriptions/{{.Lang}}/markets/{{.MarketID}}/variants/{{.Variant}}?include_mappings={{.IncludeMappings}}"
	pathFixture       = "/v1/sports/{{.Lang}}/sport_events/{{.EventURN}}/fixture.xml"
	pathPlayer        = "/v1/sports/{{.Lang}}/players/sr:player:{{.PlayerID}}/profile.xml"
	pathCompetitor    = "/v1/sports/{{.Lang}}/competitors/sr:competitor:{{.PlayerID}}/profile.xml"
	events            = "/v1/sports/{{.Lang}}/schedules/pre/schedule.xml?start={{.Start}}&limit={{.Limit}}"
	liveEvents        = "/v1/sports/{{.Lang}}/schedules/live/schedule.xml"
	pathTournaments   = "/v1/sports/{{.Lang}}/tournaments.xml"
	pathBookLiveEvent = "/v1/liveodds/booking-calendar/events/{{.EventURN}}/book"
	pathMatchStatuses = "/v1/descriptions/{{.Lang}}/match_status.xml"
)

// Markets all currently available markets for a language
func (a *API) Markets(lang uof.Lang) (uof.MarketDescriptions, []byte, error) {
	var mr marketsRsp
	raw, err := a.getAs(&mr, pathMarkets, &params{Lang: lang})
	return mr.Markets, raw, err
}

func (a *API) MarketVariant(lang uof.Lang, marketID int, variant string) (uof.MarketDescriptions, []byte, error) {
	var mr marketsRsp
	raw, err := a.getAs(&mr, pathMarketVariant, &params{Lang: lang, MarketID: marketID, Variant: variant})
	return mr.Markets, raw, err
}

// Fixture lists the fixture for a specified sport event
func (a *API) Fixture(lang uof.Lang, eventURN uof.URN) (*uof.Fixture, []byte, error) {
	var fr fixtureRsp
	raw, err := a.getAs(&fr, pathFixture, &params{Lang: lang, EventURN: eventURN})
	return &fr.Fixture, raw, err
}

func (a *API) Tournament(lang uof.Lang, eventURN uof.URN) (*uof.FixtureTournament, []byte, error) {
	var ft uof.FixtureTournament
	raw, err := a.getAs(&ft, pathFixture, &params{Lang: lang, EventURN: eventURN})
	return &ft, raw, err
}

func (a *API) Tournaments(lang uof.Lang) ([]uof.FixtureTournament, []byte, error) {
	var rsp tournamentsRsp
	raw, err := a.getAs(&rsp, pathTournaments, &params{Lang: lang})
	return rsp.Tournaments, raw, err
}

func (a *API) Player(lang uof.Lang, playerID int) (*uof.Player, []byte, error) {
	var pr playerRsp
	raw, err := a.getAs(&pr, pathPlayer, &params{Lang: lang, PlayerID: playerID})
	return &pr.Player, raw, err
}

func (a *API) Competitor(lang uof.Lang, playerID int) (*uof.CompetitorPlayer, []byte, error) {
	var cr competitorRsp
	raw, err := a.getAs(&cr, pathCompetitor, &params{Lang: lang, PlayerID: playerID})
	return &cr.Competitor, raw, err
}

type tournamentsRsp struct {
	Tournaments []uof.FixtureTournament `xml:"tournaments>tournament" json:"tournaments,omitempty" bson:"tournaments,omitempty"`
}

type marketsRsp struct {
	Markets uof.MarketDescriptions `xml:"market,omitempty" json:"markets,omitempty"`
	// unused
	// ResponseCode string   `xml:"response_code,attr,omitempty" json:"responseCode,omitempty"`
	// Location     string   `xml:"location,attr,omitempty" json:"location,omitempty"`
}

type playerRsp struct {
	Player      uof.Player `xml:"player" json:"player"`
	GeneratedAt time.Time  `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
}

type competitorRsp struct {
	Competitor  uof.CompetitorPlayer `xml:"competitor" json:"competitor"`
	GeneratedAt time.Time            `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
}

type fixtureRsp struct {
	Fixture     uof.Fixture `xml:"fixture" json:"fixture"`
	GeneratedAt time.Time   `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
}

type scheduleRsp struct {
	Fixtures    []uof.Fixture `xml:"sport_event,omitempty" json:"sportEvent,omitempty"`
	GeneratedAt time.Time     `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
}

// Fixtures gets all the fixtures with schedule before to
func (a *API) Fixtures(lang uof.Lang, to time.Time) (<-chan uof.Fixture, <-chan error) {
	errc := make(chan error, 1)
	out := make(chan uof.Fixture)
	go func() {
		defer close(out)
		defer close(errc)
		done := false

		parse := func(buf []byte) error {
			var sr scheduleRsp
			if err := xml.Unmarshal(buf, &sr); err != nil {
				return uof.Notice("unmarshal", err)
			}
			for _, f := range sr.Fixtures {
				out <- f
				if f.Scheduled.After(to) {
					done = true
				}
			}
			return nil
		}

		// first live events
		buf, err := a.get(liveEvents, &params{Lang: lang})
		if err != nil {
			errc <- err
			return
		}
		if err := parse(buf); err != nil {
			errc <- err
			return
		}

		// than all events which has scheduled before to
		limit := 1000
		for start := 0; true; start += limit {
			buf, err := a.get(events, &params{Lang: lang, Start: start, Limit: limit})
			if err != nil {
				errc <- err
				return
			}
			if err := parse(buf); err != nil {
				errc <- err
				return
			}
			if done {
				return
			}
		}
	}()

	return out, errc
}

func (a *API) BookAllLiveMatches(done map[string]bool) (int, map[string]bool, error) {
	buf, err := a.get(liveEvents, &params{Lang: uof.LangEN})
	if err != nil {
		return 0, done, err
	}
	var sr scheduleRsp
	if err := xml.Unmarshal(buf, &sr); err != nil {
		return 0, done, err
	}
	booked := 0
	for _, f := range sr.Fixtures {
		if f.Status == "ended" {
			continue
		}
		key := f.URN.String()
		if _, ok := done[key]; ok {
			continue

		}
		// fmt.Printf("%s %v %s  ", f.URN, f.Scheduled, f.Status)
		if err := a.post(pathBookLiveEvent, &params{EventURN: f.URN}); err != nil {
			var ae uof.APIError
			if errors.As(err, &ae) {
				if ur, e := ae.UOFRsp(); e == nil {
					if ur.Code == "FORBIDDEN" {
						return booked, done, err
					}
					// fmt.Printf("api error code %s, response %s\n", ur.Code, ur.Message)
					if strings.Contains(ur.Message, "The match is booked") {
						done[key] = true
						booked++
					}
					done[key] = false
					continue
				}
			}
			return booked, done, err
		}
		// fmt.Printf("OK\n")
		done[key] = true
		booked++
	}
	return booked, done, nil
}

func (a *API) MatchStatuses(lang uof.Lang) ([]MatchStatus, []byte, error) {
	var ms matchStatusesRsp
	raw, err := a.getAs(&ms, pathMatchStatuses, &params{Lang: lang})
	return ms.Stasuses, raw, err
}

type matchStatusesRsp struct {
	Stasuses []MatchStatus `xml:"match_status,omitempty"`
}

type MatchStatus struct {
	ID          int    `xml:"id,attr" json:"id"`
	Description string `xml:"description,attr" json:"description,omitempty"`
	Period      int    `xml:"period_number,attr" json:"period,omitempty"`
	Sports      []int  `json:"sports,omitempty"`
}

func (o *MatchStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MatchStatus
	var overlay struct {
		*T
		SportIds []struct {
			ID string `xml:"id,attr"`
		} `xml:"sports>sport"`
	}
	overlay.T = (*T)(o)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	if l := len(overlay.SportIds); l > 0 {
		sports := make([]int, 0, l)
		for _, s := range overlay.SportIds {
			sports = append(sports, uof.URN(s.ID).ID())
		}
		o.Sports = sports
	}
	return nil
}
