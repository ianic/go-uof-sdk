package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/minus5/go-uof-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplate(t *testing.T) {
	path := runTemplate(startScenario, &params{ScenarioID: 1, Speed: 2, MaxDelay: 3})
	assert.Equal(t, "/v1/replay/scenario/play/1?speed=2&max_delay=3&use_replay_timestamp=false", path)
}

const EnvToken = "UOF_TOKEN"

// this test depends on UOF_TOKEN environment variable
// to be set to the staging access token
// run it as:
//    UOF_TOKEN=my-token go test -v
func TestIntegration(t *testing.T) {
	token, ok := os.LookupEnv(EnvToken)
	if !ok {
		t.Skip("integration token not found")
	}

	a, err := Staging(context.TODO(), token)
	assert.NoError(t, err)

	tests := []struct {
		name string
		f    func(t *testing.T, a *API)
	}{
		{"markets", testMarkets},
		{"marketVariant", testMarketVariant},
		{"fixture", testFixture},
		{"player", testPlayer},
		{"fixtures", testFixtures},
	}
	for _, s := range tests {
		t.Run(s.name, func(t *testing.T) { s.f(t, a) })
	}
}

func TestBetCancelSeedData(t *testing.T) {
	if os.Getenv("seed_data") == "" {
		t.Skip("skipping test; $seed_data env not set")
	}
	token, ok := os.LookupEnv(EnvToken)
	if !ok {
		t.Skip("integration token not found")
	}

	a, err := Staging(context.TODO(), token)
	assert.NoError(t, err)

	mm, _, err := a.Markets(uof.LangEN)
	assert.NoError(t, err)

	buf, err := json.Marshal(mm.Groups())
	assert.NoError(t, err)
	fmt.Printf("bet cancel seed data: \n%s\n", buf)
}

func testMarkets(t *testing.T, a *API) {
	lang := uof.LangEN
	mm, _, err := a.Markets(lang)
	assert.Nil(t, err)

	assert.True(t, len(mm) >= 992)
	m := mm.Find(1)
	assert.Equal(t, "1x2", m.Name)
}

func testMarketVariant(t *testing.T, a *API) {
	lang := uof.LangEN
	mm, _, err := a.MarketVariant(lang, 241, "sr:exact_games:bestof:5")
	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Len(t, mm, 1)
	m := mm[0]
	assert.Equal(t, "Exact games", m.Name)
	assert.Len(t, m.Outcomes, 3)
}

func testFixture(t *testing.T, a *API) {
	lang := uof.LangEN
	f, _, err := a.Fixture(lang, "sr:match:8696826")
	assert.Nil(t, err)
	assert.Equal(t, "IK Oddevold", f.Home.Name)

	tf, _, err := a.Tournament(lang, "vto:season:1856707")
	assert.Nil(t, err)
	assert.NotNil(t, tf)
	//pp(tf)
}

func testPlayer(t *testing.T, a *API) {
	lang := uof.LangEN
	p, _, err := a.Player(lang, 947)
	assert.NoError(t, err)
	assert.Equal(t, "Lee Barnard", p.FullName)
}

func testFixtures(t *testing.T, a *API) {
	lang := uof.LangEN
	to := time.Now() //.Add(24*3*time.Hour)
	in, errc := a.Fixtures(lang, to)
	for f := range in {
		if testing.Verbose() {
			fmt.Printf("\t%s\n", f.PP())
			//pp(f)
		}
	}
	go func() {
		for err := range errc {
			panic(err)
		}
	}()
}

// PP prety print object
func pp(o interface{}) {
	buf, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}

func TestListLive(t *testing.T) {
	t.Skip("interactive test")

	done := make(map[string]bool)
	booked, _, err := staging(t).BookAllLiveMatches(done)
	require.NoError(t, err)
	fmt.Println("booked matches", booked)
}

func TestMatchStatus(t *testing.T) {
	t.Skip("interactive test")

	ms, _, err := staging(t).MatchStatuses(uof.LangEN)
	require.NoError(t, err)
	pp(ms)
}

func TestSports(t *testing.T) {
	t.Skip("interactive test")

	ss, _, err := staging(t).Sports(uof.LangEN)
	require.NoError(t, err)
	//fmt.Printf("%s\n", buf)
	pp(ss)
}

func staging(t *testing.T) *API {
	token, ok := os.LookupEnv(EnvToken)
	if !ok {
		t.Skip("integration token not found")
	}
	a, err := Staging(context.TODO(), token)
	require.NoError(t, err)
	return a
}

func TestTournaments(t *testing.T) {
	t.Skip("interactive test")

	ts, buf, err := staging(t).Tournaments(uof.LangEN)
	require.NoError(t, err)
	fmt.Printf("%s\n", buf)
	pp(ts)
	fmt.Printf("count %d\n", len(ts))
}

func TestSportTournaments(t *testing.T) {
	t.Skip("interactive test")

	ts, _, err := staging(t).SportTournaments(3, uof.LangEN)
	require.NoError(t, err)
	//fmt.Printf("%s\n", buf)
	pp(ts)
	fmt.Printf("count %d\n", len(ts))
}

func TestEventsForDate(t *testing.T) {
	t.Skip("interactive test")

	date := time.Now()
	ts, _, err := staging(t).EventsForDate(date, uof.LangEN)
	require.NoError(t, err)
	//fmt.Printf("%s\n", buf)
	pp(ts)
	fmt.Printf("count %d\n", len(ts))
}
