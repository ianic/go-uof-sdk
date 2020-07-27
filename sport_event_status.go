package uof

import "encoding/xml"

// The element "sport_event_status" is provided in the odds_change message.
// Status is the only required attribute for this element, and this attribute
// describes the current status of the sport-event itself (not started, live,
// ended, closed). Additional attributes are live-only attributes, and only
// provided while the match is live; additionally, which attributes are provided
// depends on the sport.
// Reference: https://docs.betradar.com/display/BD/UOF+-+Sport+event+status
type SportEventStatus struct {
	// High-level generic status of the match.
	Status EventStatus `xml:"status,attr" json:"status" bson:"status"`
	// Does Betradar have a scout watching the game.
	Reporting *EventReporting `xml:"reporting,attr,omitempty" json:"reporting,omitempty" bson:"reporting,omitempty"`
	// Current score for the home team.
	HomeScore *int `xml:"-" json:"homeScore,omitempty" bson:"homeScore,omitempty"`
	// Current score for the away team.
	AwayScore *int `xml:"-" json:"awayScore,omitempty" bson:"awayScore,omitempty"`
	// Sports-specific integer code the represents the live match status (first period, 2nd break, etc.).
	MatchStatus *int `xml:"match_status,attr" json:"matchStatus" bson:"matchStatus,omitempty"`
	// The player who has the serve at that moment.
	CurrentServer *Team `xml:"current_server,attr,omitempty" json:"currentServer,omitempty" bson:"currentServer,omitempty"`
	// The point score of the "home" player. The score will be 50 if the "home"
	// player has advantage. This attribute is also used for the tiebreak score
	// when the game is in a tiebreak.
	// (15 30 40 50)
	HomeGamescore *int `xml:"home_gamescore,attr,omitempty" json:"homeGamescore,omitempty" bson:"homeGamescore,omitempty"`
	// The point score of the "away" player. The score will be 50 if the "away"
	// player has advantage. This attribute is also used for the tiebreak score
	// when the game is in a tiebreak.
	AwayGamescore *int `xml:"away_gamescore,attr,omitempty" json:"awayGamescore,omitempty" bson:"awayGamescore,omitempty"`

	HomePenaltyScore   *int    `xml:"home_penalty_score,attr,omitempty" json:"homePenaltyScore,omitempty" bson:"homePenaltyScore,omitempty"`
	AwayPenaltyScore   *int    `xml:"away_penalty_score,attr,omitempty" json:"awayPenaltyScore,omitempty" bson:"awayPenaltyScore,omitempty"`
	HomeLegscore       *int    `xml:"home_legscore,attr,omitempty" json:"homeLegscore,omitempty" bson:"homeLegscore,omitempty"`
	AwayLegscore       *int    `xml:"away_legscore,attr,omitempty" json:"awayLegscore,omitempty" bson:"awayLegscore,omitempty"`
	ExpediteMode       *bool   `xml:"expedite_mode,attr,omitempty" json:"expediteMode,omitempty" bson:"expediteMode,omitempty"`
	Tiebreak           *bool   `xml:"tiebreak,attr,omitempty" json:"tiebreak,omitempty" bson:"tiebreak,omitempty"`
	HomeSuspend        *int    `xml:"home_suspend,attr,omitempty" json:"homeSuspend,omitempty" bson:"homeSuspend,omitempty"`
	AwaySuspend        *int    `xml:"away_suspend,attr,omitempty" json:"awaySuspend,omitempty" bson:"awaySuspend,omitempty"`
	Balls              *int    `xml:"balls,attr,omitempty" json:"balls,omitempty" bson:"balls,omitempty"`
	Strikes            *int    `xml:"strikes,attr,omitempty" json:"strikes,omitempty" bson:"strikes,omitempty"`
	Outs               *int    `xml:"outs,attr,omitempty" json:"outs,omitempty" bson:"outs,omitempty"`
	Bases              *string `xml:"bases,attr,omitempty" json:"bases,omitempty" bson:"bases,omitempty"`
	HomeBatter         *int    `xml:"home_batter,attr,omitempty" json:"homeBatter,omitempty" bson:"homeBatter,omitempty"`
	AwayBatter         *int    `xml:"away_batter,attr,omitempty" json:"awayBatter,omitempty" bson:"awayBatter,omitempty"`
	Possession         *int    `xml:"possession,attr,omitempty" json:"possession,omitempty" bson:"possession,omitempty"`
	Position           *int    `xml:"position,attr,omitempty" json:"position,omitempty" bson:"position,omitempty"`
	Try                *int    `xml:"try,attr,omitempty" json:"try,omitempty" bson:"try,omitempty"`
	Yards              *int    `xml:"yards,attr,omitempty" json:"yards,omitempty" bson:"yards,omitempty"`
	Throw              *int    `xml:"throw,attr,omitempty" json:"throw,omitempty" bson:"throw,omitempty"`
	Visit              *int    `xml:"visit,attr,omitempty" json:"visit,omitempty" bson:"visit,omitempty"`
	RemainingReds      *int    `xml:"remaining_reds,attr,omitempty" json:"remainingReds,omitempty" bson:"remainingReds,omitempty"`
	Delivery           *int    `xml:"delivery,attr,omitempty" json:"delivery,omitempty" bson:"delivery,omitempty"`
	HomeRemainingBowls *int    `xml:"home_remaining_bowls,attr,omitempty" json:"homeRemainingBowls,omitempty" bson:"homeRemainingBowls,omitempty"`
	AwayRemainingBowls *int    `xml:"away_remaining_bowls,attr,omitempty" json:"awayRemainingBowls,omitempty" bson:"awayRemainingBowls,omitempty"`
	CurrentEnd         *int    `xml:"current_end,attr,omitempty" json:"currentEnd,omitempty" bson:"currentEnd,omitempty"`
	Innings            *int    `xml:"innings,attr,omitempty" json:"innings,omitempty" bson:"innings,omitempty"`
	Over               *int    `xml:"over,attr,omitempty" json:"over,omitempty" bson:"over,omitempty"`
	HomePenaltyRuns    *int    `xml:"home_penalty_runs,attr,omitempty" json:"homePenaltyRuns,omitempty" bson:"homePenaltyRuns,omitempty"`
	AwayPenaltyRuns    *int    `xml:"away_penalty_runs,attr,omitempty" json:"awayPenaltyRuns,omitempty" bson:"awayPenaltyRuns,omitempty"`
	HomeDismissals     *int    `xml:"home_dismissals,attr,omitempty" json:"homeDismissals,omitempty" bson:"homeDismissals,omitempty"`
	AwayDismissals     *int    `xml:"away_dismissals,attr,omitempty" json:"awayDismissals,omitempty" bson:"awayDismissals,omitempty"`
	CurrentCtTeam      *Team   `xml:"current_ct_team,attr,omitempty" json:"currentCtTeam,omitempty" bson:"currentCtTeam,omitempty"`

	Clock        *Clock        `xml:"clock,omitempty" json:"clock,omitempty" bson:"clock,omitempty"`
	PeriodScores []PeriodScore `xml:"period_scores>period_score,omitempty" json:"periodScores,omitempty" bson:"periodScores,omitempty"`
	Results      []Result      `xml:"results>result,omitempty" json:"results,omitempty" bson:"results,omitempty"`
	Statistics   *Statistics   `xml:"statistics,omitempty" json:"statistics,omitempty" bson:"statistics,omitempty"`
}

func (o *SportEventStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SportEventStatus
	var overlay struct {
		*T
		HomeScoreF *float64 `xml:"home_score,attr,omitempty"`
		AwayScoreF *float64 `xml:"away_score,attr,omitempty"`
	}
	overlay.T = (*T)(o)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	if overlay.HomeScoreF != nil {
		v := int(*overlay.HomeScoreF)
		o.HomeScore = &v
	}
	if overlay.AwayScoreF != nil {
		v := int(*overlay.AwayScoreF)
		o.AwayScore = &v
	}
	return nil
}

// The sport_event_status may contain a clock element. This clock element
// includes various clock/time attributes that are sports specific.
type Clock struct {
	// The playing minute of the match (or minute:second if available)
	// mm:ss (42:10)
	MatchTime *ClockTime `xml:"match_time,attr,omitempty" json:"matchTime,omitempty" bson:"matchTime,omitempty"`
	// How far into stoppage time is the match in minutes
	// mm:ss
	StoppageTime *ClockTime `xml:"stoppage_time,attr,omitempty" json:"stoppageTime,omitempty" bson:"stoppageTime,omitempty"`
	// Set to what the announce stoppage time is announced
	// mm:ss
	StoppageTimeAnnounced *ClockTime `xml:"stoppage_time_announced,attr,omitempty" json:"stoppageTimeAnnounced,omitempty" bson:"stoppageTimeAnnounced,omitempty"`
	// How many minutes remains of the match
	// mm:ss
	RemainingTime *ClockTime `xml:"remaining_time,attr,omitempty" json:"remainingTime,omitempty" bson:"remainingTime,omitempty"`
	// How much time remains in the current period
	// mm:ss
	RemainingTimeInPeriod *ClockTime `xml:"remaining_time_in_period,attr,omitempty" json:"remainingTimeInPeriod,omitempty" bson:"remainingTimeInPeriod,omitempty"`
	// true if the match clock is stopped otherwise false
	Stopped *bool `xml:"stopped,attr,omitempty" json:"stopped,omitempty" bson:"stopped,omitempty"`
}

type PeriodScore struct {
	// The match status of an event gives an indication of which context the
	// current match is in. Complete list available at:
	// /v1/descriptions/en/match_status.xml
	MatchStatusCode *int `xml:"match_status_code,attr" json:"matchStatusCode" bson:"matchStatusCode,omitempty"`
	// Indicates what regular period this is
	Number *int `xml:"number,attr" json:"number" bson:"number,omitempty"`
	// The number of points/goals/games the competitor designated as "home" has
	// scored for this period.
	HomeScore *int `xml:"home_score,attr" json:"homeScore" bson:"homeScore,omitempty"`
	// The number of points/goals/games the competitor designated as "away" has
	// scored for this period.
	AwayScore *int `xml:"away_score,attr" json:"awayScore" bson:"awayScore,omitempty"`
}

type Result struct {
	MatchStatusCode *int `xml:"match_status_code,attr" json:"matchStatusCode" bson:"matchStatusCode,omitempty"`
	HomeScore       *int `xml:"home_score,attr" json:"homeScore" bson:"homeScore,omitempty"`
	AwayScore       *int `xml:"away_score,attr" json:"awayScore" bson:"awayScore,omitempty"`
}

type Statistics struct {
	YellowCards    *StatisticsScore `xml:"yellow_cards,omitempty" json:"yellowCards,omitempty" bson:"yellowCards,omitempty"`
	RedCards       *StatisticsScore `xml:"red_cards,omitempty" json:"redCards,omitempty" bson:"redCards,omitempty"`
	YellowRedCards *StatisticsScore `xml:"yellow_red_cards,omitempty" json:"yellowRedCards,omitempty" bson:"yellowRedCards,omitempty"`
	Corners        *StatisticsScore `xml:"corners,omitempty" json:"corners,omitempty" bson:"corners,omitempty"`
}

type StatisticsScore struct {
	Home int `xml:"home,attr" json:"home" bson:"home,omitempty"`
	Away int `xml:"away,attr" json:"away" bson:"away,omitempty"`
}
