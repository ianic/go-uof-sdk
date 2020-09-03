package uof

import (
	"encoding/xml"
	"fmt"
	"time"
)

type FixtureRsp struct {
	Fixture     Fixture   `xml:"fixture" json:"fixture" bson:"fixture,omitempty"`
	GeneratedAt time.Time `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty" bson:"generatedAt,omitempty"`
}

// Fixtures describe static or semi-static information about matches and races.
// Reference: https://docs.betradar.com/display/BD/UOF+-+Fixtures+in+the+API
type Fixture struct {
	ID                 int       `xml:"-" json:"id" bson:"id,omitempty"`
	URN                URN       `xml:"id,attr,omitempty" json:"urn" bson:"urn,omitempty"`
	StartTime          time.Time `xml:"start_time,attr,omitempty" json:"startTime,omitempty" bson:"startTime,omitempty"`
	StartTimeConfirmed bool      `xml:"start_time_confirmed,attr,omitempty" json:"startTimeConfirmed,omitempty" bson:"startTimeConfirmed,omitempty"`
	StartTimeTbd       bool      `xml:"start_time_tbd,attr,omitempty" json:"startTimeTbd,omitempty" bson:"startTimeTbd,omitempty"`
	NextLiveTime       time.Time `xml:"next_live_time,attr,omitempty" json:"nextLiveTime,omitempty" bson:"nextLiveTime,omitempty"`
	Liveodds           string    `xml:"liveodds,attr,omitempty" json:"liveodds,omitempty" bson:"liveodds,omitempty"`
	Status             string    `xml:"status,attr,omitempty" json:"status,omitempty" bson:"status,omitempty"`
	Name               string    `xml:"name,attr,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	Type               string    `xml:"type,attr,omitempty" json:"type,omitempty" bson:"type,omitempty"`
	Scheduled          time.Time `xml:"scheduled,attr,omitempty" json:"scheduled,omitempty" bson:"scheduled,omitempty"`
	ScheduledEnd       time.Time `xml:"scheduled_end,attr,omitempty" json:"scheduledEnd,omitempty" bson:"scheduledEnd,omitempty"`
	ReplacedBy         string    `xml:"replaced_by,attr,omitempty" json:"replacedBy,omitempty" bson:"replacedBy,omitempty"`

	Sport      Sport      `xml:"sport" json:"sport" bson:"sport,omitempty"`
	Category   Category   `xml:"category" json:"category" bson:"category,omitempty"`
	Tournament Tournament `xml:"tournament,omitempty" json:"tournament,omitempty" bson:"tournament,omitempty"`

	Round  Round  `xml:"tournament_round,omitempty" json:"round,omitempty" bson:"round,omitempty"`
	Season Season `xml:"season,omitempty" json:"season,omitempty" bson:"season,omitempty"`
	Venue  Venue  `xml:"venue,omitempty" json:"venue,omitempty" bson:"venue,omitempty"`

	ProductInfo ProductInfo  `xml:"product_info,omitempty" json:"productInfo,omitempty" bson:"productInfo,omitempty"`
	Competitors []Competitor `xml:"competitors>competitor,omitempty" json:"competitors,omitempty" bson:"competitors,omitempty"`
	TvChannels  []TvChannel  `xml:"tv_channels>tv_channel,omitempty" json:"tvChannels,omitempty" bson:"tvChannels,omitempty"`

	Home Competitor `json:"home" bson:"home,omitempty"`
	Away Competitor `json:"away" bson:"away,omitempty"`

	ExtraInfo []ExtraInfo  `xml:"extra_info>info,omitempty" json:"extraInfo,omitempty" bson:"extraInfo,omitempty"`
	Races     []SportEvent `xml:"races>sport_event,omitempty" json:"races,omitempty" bson:"races,omitempty"`
	// this also exists but we are skiping for the time being
	//ReferenceIDs         ReferenceIDs         `xml:"reference_ids,omitempty" json:"referenceId`s,omitempty"`
	//SportEventConditions SportEventConditions `xml:"sport_event_conditions,omitempty" json:"sportEventConditions,omitempty"`
	//DelayedInfo DelayedInfo `xml:"delayed_info,omitempty" json:"delayedInfo,omitempty"`
	//CoverageInfo CoverageInfo `xml:"coverage_info,omitempty" json:"coverageInfo,omitempty"`
	//ScheduledStartTimeChanges []ScheduledStartTimeChange `xml:"scheduled_start_time_changes>scheduled_start_time_change,omitempty" json:"scheduledStartTimeChanges,omitempty"`
	//Parent *ParentStage `xml:"parent,omitempty" json:"parent,omitempty"`
}

type FixtureTournament struct {
	ID         int        `xml:"-" json:"id" bson:"id,omitempty"`
	URN        URN        `xml:"id,attr,omitempty" json:"urn" bson:"urn,omitempty"`
	Name       string     `xml:"name,attr,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	Sport      Sport      `xml:"sport" json:"sport" bson:"sport,omitempty"`
	Category   Category   `xml:"category" json:"category" bson:"category,omitempty"`
	Tournament Tournament `xml:"tournament,omitempty" json:"tournament,omitempty" bson:"tournament,omitempty"`
	Season     Season     `xml:"season,omitempty" json:"season,omitempty" bson:"season,omitempty"`
	Groups     []Group    `xml:"groups>group,omitempty" json:"groups,omitempty" bson:"groups,omitempty"`
}

type Group struct {
	Name        string       `xml:"name,attr,omitempty" json:"name" bson:"name,omitempty"`
	Competitors []Competitor `xml:"competitor,omitempty" json:"competitors,omitempty" bson:"competitors,omitempty"`
}

type Tournament struct {
	ID   int    `json:"id" bson:"id,omitempty"`
	Name string `xml:"name,attr" json:"name" bson:"name,omitempty"`
}

type Sport struct {
	ID   int    `json:"id" bson:"id,omitempty"`
	Name string `xml:"name,attr" json:"name" bson:"name,omitempty"`
}

type Category struct {
	ID          int    `json:"id" bson:"id,omitempty"`
	Name        string `xml:"name,attr" json:"name" bson:"name,omitempty"`
	CountryCode string `xml:"country_code,attr,omitempty" json:"countryCode,omitempty" bson:"countryCode,omitempty"`
}

type Competitor struct {
	ID           int                `json:"id" bson:"id,omitempty"`
	Qualifier    string             `xml:"qualifier,attr,omitempty" json:"qualifier,omitempty" bson:"qualifier,omitempty"`
	Name         string             `xml:"name,attr" json:"name" bson:"name,omitempty"`
	Abbreviation string             `xml:"abbreviation,attr" json:"abbreviation" bson:"abbreviation,omitempty"`
	Country      string             `xml:"country,attr,omitempty" json:"country,omitempty" bson:"country,omitempty"`
	CountryCode  string             `xml:"country_code,attr,omitempty" json:"countryCode,omitempty" bson:"countryCode,omitempty"`
	Virtual      bool               `xml:"virtual,attr,omitempty" json:"virtual,omitempty" bson:"virtual,omitempty"`
	Players      []CompetitorPlayer `xml:"players>player,omitempty" json:"players,omitempty" bson:"players,omitempty"`
	//ReferenceIDs CompetitorReferenceIDs `xml:"reference_ids,omitempty" json:"referenceIds,omitempty"`
}

type CompetitorPlayer struct {
	ID           int    `json:"id" bson:"id,omitempty"`
	Name         string `xml:"name,attr" json:"name" bson:"name,omitempty"`
	Abbreviation string `xml:"abbreviation,attr" json:"abbreviation" bson:"abbreviation,omitempty"`
	Nationality  string `xml:"nationality,attr,omitempty" json:"nationality,omitempty" bson:"nationality,omitempty"`
}

type Venue struct {
	ID             int    `json:"id" bson:"id,omitempty"`
	Name           string `xml:"name,attr" json:"name" bson:"name,omitempty"`
	Capacity       int    `xml:"capacity,attr,omitempty" json:"capacity,omitempty" bson:"capacity,omitempty"`
	CityName       string `xml:"city_name,attr,omitempty" json:"cityName,omitempty" bson:"cityName,omitempty"`
	CountryName    string `xml:"country_name,attr,omitempty" json:"countryName,omitempty" bson:"countryName,omitempty"`
	CountryCode    string `xml:"country_code,attr,omitempty" json:"countryCode,omitempty" bson:"countryCode,omitempty"`
	MapCoordinates string `xml:"map_coordinates,attr,omitempty" json:"mapCoordinates,omitempty" bson:"mapCoordinates,omitempty"`
}

type TvChannel struct {
	Name string `xml:"name,attr" json:"name" bson:"name,omitempty"`
	// seams to be always zero
	// StartTime time.Time `xml:"start_time,attr,omitempty" json:"startTime,omitempty"`
}

type StreamingChannel struct {
	ID   int    `xml:"id,attr" json:"id" bson:"id,omitempty"`
	Name string `xml:"name,attr" json:"name" bson:"name,omitempty"`
}
type ProductInfoLink struct {
	Name string `xml:"name,attr" json:"name" bson:"name,omitempty"`
	Ref  string `xml:"ref,attr" json:"ref" bson:"ref,omitempty"`
}

type Round struct {
	ID                  int    `xml:"betradar_id,attr,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	Type                string `xml:"type,attr,omitempty" json:"type,omitempty" bson:"type,omitempty"`
	Number              int    `xml:"number,attr,omitempty" json:"number,omitempty" bson:"number,omitempty"`
	Name                string `xml:"name,attr,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	GroupLongName       string `xml:"group_long_name,attr,omitempty" json:"groupLongName,omitempty" bson:"groupLongName,omitempty"`
	Group               string `xml:"group,attr,omitempty" json:"group,omitempty" bson:"group,omitempty"`
	GroupID             string `xml:"group_id,attr,omitempty" json:"groupId,omitempty" bson:"groupId,omitempty"`
	CupRoundMatches     int    `xml:"cup_round_matches,attr,omitempty" json:"cupRoundMatches,omitempty" bson:"cupRoundMatches,omitempty"`
	CupRoundMatchNumber int    `xml:"cup_round_match_number,attr,omitempty" json:"cupRoundMatchNumber,omitempty" bson:"cupRoundMatchNumber,omitempty"`
	OtherMatchID        string `xml:"other_match_id,attr,omitempty" json:"otherMatchId,omitempty" bson:"otherMatchId,omitempty"`
}

type Season struct {
	ID        int    `json:"id" bson:"id,omitempty"`
	StartDate string `xml:"start_date,attr" json:"startDate" bson:"startDate,omitempty"`
	EndDate   string `xml:"end_date,attr" json:"endDate" bson:"endDate,omitempty"`
	StartTime string `xml:"start_time,attr,omitempty" json:"startTime,omitempty" bson:"startTime,omitempty"`
	EndTime   string `xml:"end_time,attr,omitempty" json:"endTime,omitempty" bson:"endTime,omitempty"`
	Year      string `xml:"year,attr,omitempty" json:"year,omitempty" bson:"year,omitempty"`
	Name      string `xml:"name,attr" json:"name" bson:"name,omitempty"`
	//TournamentID string    `xml:"tournament_id,attr,omitempty" json:"tournamentId,omitempty"`
}

// type ParentStage struct {
// 	URN          URN       `xml:"id,attr,omitempty" json:"urn,omitempty"`
// 	Name         string    `xml:"name,attr,omitempty" json:"name,omitempty"`
// 	Type         string    `xml:"type,attr,omitempty" json:"type,omitempty"`
// 	Scheduled    time.Time `xml:"scheduled,attr,omitempty" json:"scheduled,omitempty"`
// 	StartTimeTbd bool      `xml:"start_time_tbd,attr,omitempty" json:"startTimeTbd,omitempty"`
// 	ScheduledEnd time.Time `xml:"scheduled_end,attr,omitempty" json:"scheduledEnd,omitempty"`
// 	ReplacedBy   string    `xml:"replaced_by,attr,omitempty" json:"replacedBy,omitempty"`
// }

// type ScheduledStartTimeChange struct {
// 	OldTime   time.Time `xml:"old_time,attr" json:"oldTime"`
// 	NewTime   time.Time `xml:"new_time,attr" json:"newTime"`
// 	ChangedAt time.Time `xml:"changed_at,attr" json:"changedAt"`
// }

type ProductInfo struct {
	Streaming            []StreamingChannel `xml:"streaming>channel,omitempty" json:"streaming,omitempty" bson:"streaming,omitempty"`
	IsInLiveScore        string             `xml:"is_in_live_score,omitempty" json:"isInLiveScore,omitempty" bson:"isInLiveScore,omitempty"`
	IsInHostedStatistics string             `xml:"is_in_hosted_statistics,omitempty" json:"isInHostedStatistics,omitempty" bson:"isInHostedStatistics,omitempty"`
	IsInLiveCenterSoccer string             `xml:"is_in_live_center_soccer,omitempty" json:"isInLiveCenterSoccer,omitempty" bson:"isInLiveCenterSoccer,omitempty"`
	IsAutoTraded         string             `xml:"is_auto_traded,omitempty" json:"isAutoTraded,omitempty" bson:"isAutoTraded,omitempty"`
	Links                []ProductInfoLink  `xml:"links>link,omitempty" json:"links,omitempty" bson:"links,omitempty"`
}

// ExtraInfo covers additional fixture information about the match,
// such as coverage information, extended markets offer, additional rules etc.
type ExtraInfo struct {
	Key   string `xml:"key,attr,omitempty" json:"key,omitempty" bson:"key,omitempty"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty" bson:"value,omitempty"`
}

// SportEvent covers information about scheduled races in a stage
// For VHC and VDR information is in vdr/vhc:stage:<int> fixture with type="parent"
type SportEvent struct {
	ID           string    `xml:"id,attr,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	Name         string    `xml:"name,attr,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	Type         string    `xml:"type,attr,omitempty" json:"type,omitempty" bson:"type,omitempty"`
	Scheduled    time.Time `xml:"scheduled,attr,omitempty" json:"scheduled,omitempty" bson:"scheduled,omitempty"`
	ScheduledEnd time.Time `xml:"scheduled_end,attr,omitempty" json:"scheduled_end,omitempty" bson:"scheduledEnd,omitempty"`
}

func (f *Fixture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Fixture
	var overlay struct {
		*T
		Tournament *struct {
			URN      URN      `xml:"id,attr"`
			Name     string   `xml:"name,attr"`
			Sport    Sport    `xml:"sport"`
			Category Category `xml:"category"`
		} `xml:"tournament,omitempty"`
	}
	overlay.T = (*T)(f)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	f.ID = overlay.URN.EventID()
	f.Sport = overlay.Tournament.Sport
	f.Category = overlay.Tournament.Category
	f.Tournament.ID = overlay.Tournament.URN.ID()
	f.Tournament.Name = overlay.Tournament.Name

	for _, c := range f.Competitors {
		if c.Qualifier == "home" {
			f.Home = c
		}
		if c.Qualifier == "away" {
			f.Away = c
		}
	}
	return nil
}

func (t *Sport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Sport
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *Category) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Category
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *Season) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Season
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *Venue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Venue
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *Competitor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Competitor
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *CompetitorPlayer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompetitorPlayer
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *FixtureTournament) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FixtureTournament
	var overlay struct {
		*T
		CurrentSeason Season `xml:"current_season"`
		Tournament    *struct {
			URN      URN      `xml:"id,attr"`
			Name     string   `xml:"name,attr"`
			Sport    Sport    `xml:"sport"`
			Category Category `xml:"category"`
			Season   Season   `xml:"current_season"`
		} `xml:"tournament,omitempty"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	if overlay.Tournament != nil {
		t.Sport = overlay.Tournament.Sport
		t.Category = overlay.Tournament.Category
		if !overlay.Tournament.URN.Empty() {
			t.Tournament.ID = overlay.Tournament.URN.ID()
			t.Tournament.Name = overlay.Tournament.Name
			t.ID = overlay.Tournament.URN.ID()
			t.URN = overlay.Tournament.URN
		}
	}
	if t.Season.ID == 0 && overlay.CurrentSeason.ID != 0 {
		t.Season = overlay.CurrentSeason
	}

	t.Tournament.ID = t.URN.ID()
	t.Tournament.Name = t.Name
	t.ID = t.URN.ID()

	return nil
}

// PP pretty prints fixure row
func (f *Fixture) PP() string {
	name := fmt.Sprintf("%s - %s", f.Home.Name, f.Away.Name)
	return fmt.Sprintf("%-90s %12s %15s", name, f.Scheduled.Format("02.01. 15:04"), f.Status)
}
