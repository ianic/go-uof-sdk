package uof

import "encoding/xml"

type BetSettlement struct {
	EventID   int      `json:"eventId" bson:"eventId,omitempty"`
	EventURN  URN      `xml:"event_id,attr" json:"eventURN" bson:"eventURN,omitempty"`
	Producer  Producer `xml:"product,attr" json:"producer" bson:"producer,omitempty"`
	Timestamp int      `xml:"timestamp,attr" json:"timestamp" bson:"timestamp,omitempty"`
	RequestID *int     `xml:"request_id,attr,omitempty" json:"requestId,omitempty" bson:"requestId,omitempty"`
	// Is this bet-settlement sent as a consequence of scouts reporting the
	// results live (1) or is this bet-settlement sent post-match when the
	// official results have been confirmed (2)
	Certainty *int8                 `xml:"certainty,attr" json:"certainty" bson:"certainty,omitempty"` // May be one of 1, 2
	Markets   []BetSettlementMarket `xml:"outcomes>market" json:"markets" bson:"markets,omitempty"`
}

type BetSettlementMarket struct {
	ID         int               `xml:"id,attr" json:"id" bson:"id,omitempty"`
	LineID     int               `json:"lineId" bson:"lineId,omitempty"`
	Specifiers map[string]string `json:"specifiers,omitempty" bson:"specifiers,omitempty"`
	// Describes the reason for voiding certain outcomes for a particular market.
	// Only set if at least one of the outcomes have a void_factor. A list of void
	// reasons can be found above this table or by using the API at
	// https://iodocs.betradar.com/unifiedfeed#Betting-descriptions-GET-Void-reasons.
	VoidReason *int                   `xml:"void_reason,attr,omitempty" json:"voidReason,omitempty" bson:"voidReason,omitempty"`
	Result     *string                `xml:"result,attr,omitempty" json:"result,omitempty" bson:"result,omitempty"`
	Outcomes   []BetSettlementOutcome `xml:"outcome" json:"outcomes" bson:"outcomes,omitempty"`
}

type BetSettlementOutcome struct {
	ID             int           `json:"id" bson:"id"`
	PlayerID       int           `json:"playerId,omitempty" bson:"playerId,omitempty"`
	Result         OutcomeResult `json:"result" bson:"result"`
	DeadHeatFactor float64       `json:"deadHeatFactor,omitempty" bson:"deadHeatFactor,omitempty"`
}

type RollbackBetSettlement struct {
	EventID   int               `json:"eventId" bson:"eventId,omitempty"`
	EventURN  URN               `xml:"event_id,attr" json:"eventURN" bson:"eventURN,omitempty"`
	Producer  Producer          `xml:"product,attr" json:"producer" bson:"producer,omitempty"`
	Timestamp int               `xml:"timestamp,attr" json:"timestamp" bson:"timestamp,omitempty"`
	RequestID *int              `xml:"request_id,attr,omitempty" json:"requestId,omitempty" bson:"requestId,omitempty"`
	Markets   []BetCancelMarket `xml:"market" json:"markets" bson:"markets,omitempty"`
}

func (t *BetSettlement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BetSettlement
	var overlay struct {
		*T
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.EventID = t.EventURN.EventID()
	return nil
}

func (t *RollbackBetSettlement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RollbackBetSettlement
	var overlay struct {
		*T
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.EventID = t.EventURN.EventID()
	return nil
}

func (t *BetSettlementMarket) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BetSettlementMarket
	var overlay struct {
		*T
		Specifiers         string `xml:"specifiers,attr,omitempty"`
		ExtendedSpecifiers string `xml:"extended_specifiers,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.Specifiers = toSpecifiers(overlay.Specifiers, overlay.ExtendedSpecifiers)
	t.LineID = toLineID(overlay.Specifiers)
	return nil
}

func (t *BetSettlementOutcome) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BetSettlementOutcome
	var overlay struct {
		*T
		ID             string   `xml:"id,attr"`
		Result         *int     `xml:"result,attr"`
		VoidFactor     *float64 `xml:"void_factor,attr,omitempty"`
		DeadHeatFactor *float64 `xml:"dead_heat_factor,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = toOutcomeID(overlay.ID)
	t.PlayerID = toPlayerID(overlay.ID)
	t.Result = toResult(overlay.Result, overlay.VoidFactor, overlay.DeadHeatFactor)
	if t.Result == OutcomeResultWinWithDeadHead && overlay.DeadHeatFactor != nil {
		t.DeadHeatFactor = *overlay.DeadHeatFactor
	}
	return nil
}

//The following list includes all possible combinations of outcome (result) and void_factor:
//  result="0" and no void_factor: Lose entire bet
//  result="1" and no void_factor: Win entire bet
//  result="0" and void_factor="1": Refund entire bet
//  result="1" and void_factor="0.5": Refund half bet and win other half
//  result="0" and void_factor="0.5": Refund half bet and lose other half.
// If the bet on an outcome should be refunded completely void-factor is set to
// 1.0. If half of the bet on an outcome should be refunded void_factor is set
// to 0.5.
// Reference: https://docs.betradar.com/display/BD/UOF+-+Bet+settlement
func toResult(resultN *int, voidFactorN *float64, deadHeatFactor *float64) OutcomeResult {
	if resultN == nil {
		return OutcomeResultUnknown
	}
	result := *resultN
	voidFactor := float64(0)
	if voidFactorN != nil {
		voidFactor = *voidFactorN
	}

	if result == -1 {
		return OutcomeResultUnknown
	}
	if result == 0 && voidFactor == 0 {
		return OutcomeResultLose
	}
	if result == 1 && voidFactor == 0 {
		if deadHeatFactor != nil && *deadHeatFactor > 0 {
			return OutcomeResultWinWithDeadHead
		}
		return OutcomeResultWin
	}
	if result == 0 && voidFactor == 1 {
		return OutcomeResultVoid
	}
	if result == 0 && voidFactor == 0.5 {
		return OutcomeResultHalfLose
	}
	if result == 1 && voidFactor == 0.5 {
		return OutcomeResultHalfWin
	}
	return OutcomeResultUnknown
}
