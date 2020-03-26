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
	producer := m.Producer.Code()
	if m.Producer.Sports() {
		producer = "sport"
	}

	switch m.Type.Kind() {
	case uof.MessageKindEvent:
		if m.Type == uof.MessageTypeOddsChange {
			return fmt.Sprintf("/log/events/%s/%d/%13d", producer, m.EventID, m.ReceivedAt)
		}
		return fmt.Sprintf("/log/events/%s/%d/%13d-%s", producer, m.EventID, m.ReceivedAt, m.Type)
	case uof.MessageKindLexicon:
		switch m.Type {
		case uof.MessageTypePlayer:
			return fmt.Sprintf("/state/%s/%s/players/%08d/%13d", producer, m.Lang, m.Player.ID, m.ReceivedAt)
		case uof.MessageTypeMarkets:
			if len(m.Markets) > 1 {
				return fmt.Sprintf("/state/%s/%s/markets/%s/%13d", producer, m.Lang, m.Lang, m.ReceivedAt)
			}
			s := m.Markets[0]
			return fmt.Sprintf("/state/%s/%s/markets/%08d-%08d/%13d", producer, m.Lang, s.ID, s.VariantID, m.ReceivedAt)
		case uof.MessageTypeFixture:
			if m.EventURN == "" {
				return fmt.Sprintf("/state/%s/%s/fixtures/%08d/%13d", producer, m.Lang, m.EventID, m.ReceivedAt)
			}
			return fmt.Sprintf("/state/%s/%s/fixtures/%d/%13d", producer, m.Lang, m.EventID, m.ReceivedAt)
		case uof.MessageTypeCompetitor:
			return fmt.Sprintf("/state/%s/%s/competitors/%d/%13d", producer, m.Lang, m.Competitor.ID, m.ReceivedAt)
		case uof.MessageTypeTournament:
			return fmt.Sprintf("/state/%s/%s/tournaments/%d/%13d", producer, m.Lang, m.EventID, m.ReceivedAt)
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
