package uof

import "encoding/xml"

// A bet_cancel message is sent when a bet made on a particular market needs to
// be cancelled and refunded due to an error (which is different to a
// bet-settlement/refund).
// Reference: https://docs.betradar.com/display/BD/UOF+-+Bet+cancel
type BetCancel struct {
	EventID   int      `json:"eventId" bson:"eventId,omitempty"`
	EventURN  URN      `xml:"event_id,attr" json:"eventURN" bson:"eventURN,omitempty"`
	Producer  Producer `xml:"product,attr" json:"producer" bson:"producer,omitempty"`
	Timestamp int      `xml:"timestamp,attr" json:"timestamp" bson:"timestamp,omitempty"`
	RequestID *int     `xml:"request_id,attr,omitempty" json:"requestId,omitempty" bson:"requestId,omitempty"`
	// If start and end time are specified, they designate a range in time for
	// which bets made should be cancelled. If there is an end_time but no
	// start_time, this means cancel all bets placed before the specified time. If
	// there is a start_time but no end_time this means, cancel all bets placed
	// after the specified start_time.
	StartTime    *int              `xml:"start_time,attr,omitempty" json:"startTime,omitempty" bson:"startTime,omitempty"`
	EndTime      *int              `xml:"end_time,attr,omitempty" json:"endTime,omitempty" bson:"endTime,omitempty"`
	SupercededBy *string           `xml:"superceded_by,attr,omitempty" json:"supercededBy,omitempty" bson:"supercededBy,omitempty"`
	Markets      []BetCancelMarket `xml:"market" json:"markets" bson:"markets,omitempty"`
}

type BetCancelMarket struct {
	ID         int               `xml:"id,attr" json:"id" bson:"id,omitempty"`
	LineID     int               `json:"lineId" bson:"lineId,omitempty"`
	VariantID  int               `json:"variantId,omitempty" bson:"variantId,omitempty"`
	Specifiers map[string]string `json:"specifiers,omitempty" bson:"specifiers,omitempty"`
	VoidReason *int              `xml:"void_reason,attr,omitempty" json:"voidReason,omitempty" bson:"voidReason,omitempty"`
}

// A Rollback_bet_cancel message is sent when a previous bet cancel should be
// undone (if possible). This may happen, for example, if a Betradar operator
// mistakenly cancels the wrong market (resulting in a bet_cancel being sent)
// during the game; before realizing the mistake.
type RollbackBetCancel struct {
	EventID   int               `json:"eventId" bson:"eventId,omitempty"`
	EventURN  URN               `xml:"event_id,attr" json:"eventURN" bson:"eventURN,omitempty"`
	Producer  Producer          `xml:"product,attr" json:"producer" bson:"producer,omitempty"`
	Timestamp int               `xml:"timestamp,attr" json:"timestamp" bson:"timestamp,omitempty"`
	RequestID *int              `xml:"request_id,attr,omitempty" json:"requestId,omitempty" bson:"requestId,omitempty"`
	StartTime *int              `xml:"start_time,attr,omitempty" json:"startTime,omitempty" bson:"startTime,omitempty"`
	EndTime   *int              `xml:"end_time,attr,omitempty" json:"endTime,omitempty" bson:"endTime,omitempty"`
	Markets   []BetCancelMarket `xml:"market" json:"markets" bson:"markets,omitempty"`
}

// UnmarshalXML
func (t *BetCancel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BetCancel
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

// UnmarshalXML
func (t *BetCancelMarket) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BetCancelMarket
	var overlay struct {
		*T
		Specifiers         string `xml:"specifiers,attr,omitempty"`
		ExtendedSpecifiers string `xml:"extended_specifiers,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.Specifiers, t.LineID = toSpecifiersLineID(overlay.Specifiers, overlay.ExtendedSpecifiers)
	t.VariantID = toVariantID(variantSpecifier(t.Specifiers))
	return nil
}

// UnmarshalXML
func (t *RollbackBetCancel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RollbackBetCancel
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
