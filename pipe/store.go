package pipe

import (
	"fmt"
	"io/ioutil"
	"path"

	"os"
	"sync"

	"github.com/minus5/go-uof-sdk"
)

func InnerFileStore(root string) InnerStage {
	return Stage(func(in <-chan *uof.Message, out chan<- *uof.Message, errc chan<- error) {
		var wg sync.WaitGroup
		for m := range in {
			out <- m
			wg.Add(1)
			go func(m *uof.Message) {
				fn := root + "/" + filename(m)
				if err := save(fn, m.Marshal()); err != nil {
					errc <- uof.Notice("file save", err)
				}
				wg.Done()
			}(m)
		}
		wg.Wait()
	})
}

func FileStore(root string) ConsumerStage {
	return func(in <-chan *uof.Message) error {
		for m := range in {
			fn := root + "/" + filename(m)
			if err := save(fn, m.MarshalPretty()); err != nil {
				return err
			}
		}
		return nil
	}
}

// filename returns unique filename for the message
func filename(m *uof.Message) string {
	switch m.Type.Kind() {
	case uof.MessageKindEvent:
		producer := m.Producer.Code()
		if m.Producer.Sports() {
			producer = "sport"
		}
		if m.Type == uof.MessageTypeOddsChange {
			return fmt.Sprintf("/log/events/%s/%s/%13d", producer, m.EventURN, m.ReceivedAt)
		}
		return fmt.Sprintf("/log/events/%s/%s/%13d-%s", producer, m.EventURN, m.ReceivedAt, m.Type)
	case uof.MessageKindLexicon:
		switch m.Type {
		case uof.MessageTypePlayer:
			return fmt.Sprintf("/state/%s/players/%08d", m.Lang, m.Player.ID)
		case uof.MessageTypeMarkets:
			if len(m.Markets) > 1 {
				return fmt.Sprintf("/state/%s/markets/%s", m.Lang, m.Lang)
			}
			s := m.Markets[0]
			return fmt.Sprintf("/state/%s/markets/%08d-%08d", m.Lang, s.ID, s.VariantID)
		case uof.MessageTypeFixture:
			if m.EventURN == "" {
				return fmt.Sprintf("/state/%s/fixtures/%08d", m.Lang, m.EventID)
			}
			return fmt.Sprintf("/state/%s/fixtures/%s", m.Lang, m.EventURN)
		case uof.MessageTypeCompetitor:
			return fmt.Sprintf("/state/%s/competitors/%08d", m.Lang, m.Competitor.ID)
		case uof.MessageTypeTournament:
			return fmt.Sprintf("/state/%s/tournaments/%s", m.Lang, m.EventURN)
		}
	case uof.MessageKindSystem:
		return fmt.Sprintf("log/system/%13d-%s/%13d", m.ReceivedAt, m.Type, m.ReceivedAt)
	}
	return fmt.Sprintf("/other/%13d-%s", m.ReceivedAt, m.Type)
}

func save(filename string, buf []byte) error {
	dir, _ := path.Split(filename)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf, 0644)
}
