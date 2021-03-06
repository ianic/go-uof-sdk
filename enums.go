package uof

// The default value is active if status is not present.
type MarketStatus int8

// Reference: https://docs.betradar.com/display/BD/UOF+-+Market+status
const (
	// Active/suspended/inactive could be sent in odds change message:

	// Odds are provided and you can accept bets on the market.
	MarketStatusActive MarketStatus = 1
	// Odds continue to be provided but you should not accept bets on the market
	// for a short time (e.g. from right before a goal and until the goal has been
	// observed/confirmed).
	MarketStatusSuspended MarketStatus = -1
	// Odds are no longer provided for this market. A market can go back to Active
	// again i.e.: A total 3.5 market is deactivated since 0.5, 1.5 or 2.5 is the
	// most balanced market. However, if a goal is scored, then the 3.5 market
	// becomes the most balanced again, changing status to active. There are
	// numerous other reasons for this change as well, and it happens on a regular
	// basis.
	MarketStatusInactive MarketStatus = 0

	// During recovery the following additional status may also be sent:

	// Not a real market status. This status is normally seen under recovery, and
	// is a signal that the producer that sends this message is no longer sending
	// odds for this market. Odds will come from another producer going forward
	// (and might already have started coming from the new producer). Handed over
	// is also sent by the prematch producer when the Live Odds producer takes
	// over a market. If you have not received the live odds change yet, the
	// market should be suspended, otherwise the message can be ignored. If the
	// live odds change does not eventually appear, the market should likely be
	// deactivated.
	MarketStatusHandedOver MarketStatus = -2
	// Bet Settlement messages have been sent for this market, no further odds
	// will be provided. However, it should be noted that in rare cases (error
	// conditions), a settled market may be moved to cancelled by a bet_cancel
	// message.
	MarketStatusSettled MarketStatus = -3
	// This market has been cancelled. No further odds will be provided for this
	// market. This state is only seen during recovery for matches where the
	// system has sent out a cancellation message for that particular market.
	MarketStatusCancelled MarketStatus = -4
)

func (m MarketStatus) Val() int8 {
	return int8(m)
}

func (m *MarketStatus) PtrVal() *int8 {
	if m == nil {
		return nil
	}
	v := int8(*m)
	return &v
}

type CashoutStatus int8

const (
	// available for cashout
	CashoutStatusAvailable CashoutStatus = 1
	// temporarily unavailable for cashout
	CashoutStatusUnavailable CashoutStatus = -1
	// permanently unavailable for cashout
	CashoutStatusClosed CashoutStatus = -2
)

func (s *CashoutStatus) PtrVal() *int8 {
	if s == nil {
		return nil
	}
	v := int8(*s)
	return &v
}

type Team int8

const (
	TeamHome Team = 1
	TeamAway Team = 2
)

func (t *Team) PtrVal() *int8 {
	if t == nil {
		return nil
	}
	v := int8(*t)
	return &v
}

// Reference: https://docs.betradar.com/display/BD/UOF+-+Live+information+and+resulting
type EventStatus int8

const (
	EventStatusNotStarted EventStatus = 0 // The match has not started yet. (Alternatively, Betradar has no live coverage of the event, the match has started but we do not know this. The match will then move to closed when Betradar enters the match results)
	EventStatusLive       EventStatus = 1 // The match is live
	EventStatusSuspended  EventStatus = 2 // Used by the Premium Cricket odds producer
	EventStatusEnded      EventStatus = 3 // The match has finished, but results have not been confirmed yet.
	EventStatusClosed     EventStatus = 4 // The match is finished, results confirmed, and no more changes are expected to the results (only for events covered by pre-match producer).
	// Only one of the above statuses are possible in the odds_change message in
	// the feed. However please note that other states are available in the API,
	// but will not appear in the odds_change message. These are as following:
	EventStatusCancelled   EventStatus = 5 // The sport event (either the actual match, or this Betradar representation of the match) has been cancelled
	EventStatusDelayed     EventStatus = 6 // The sport event start has been delayed from scheduled start (most often seen for tennis).
	EventStatusInterrupted EventStatus = 7 // The sport event looks to be interrupted for a longer period than a few minutes
	EventStatusPostponed   EventStatus = 8 // The sport event has been postponed and will be played at a later date. Typically, if the later date is more than 3 days away. This sport event id will be cancelled and replaced by a new id. If the match is postponed to just one or two days from now, the same sport-id will change its state just before match start.
	EventStatusAbandoned   EventStatus = 9 // Used to indicate that Betradar has no live coverage or has lost live coverage but match is still likely ongoing.
)

func (s EventStatus) Val() int8 {
	return int8(s)
}

func (s *EventStatus) PtrVal() *int8 {
	if s == nil {
		return nil
	}
	v := int8(*s)
	return &v
}

type OutcomeResult int8

const (
	OutcomeResultUnknown         OutcomeResult = 0
	OutcomeResultLose            OutcomeResult = 1
	OutcomeResultWin             OutcomeResult = 2
	OutcomeResultVoid            OutcomeResult = 3
	OutcomeResultHalfLose        OutcomeResult = 4
	OutcomeResultHalfWin         OutcomeResult = 5
	OutcomeResultWinWithDeadHead OutcomeResult = 6
)

type OutcomeType int8

const (
	OutcomeTypeDefault OutcomeType = iota
	OutcomeTypePlayer
	OutcomeTypeCompetitor
	OutcomeTypeCompetitors
	OutcomeTypeFreeText
	OutcomeTypeUnknown OutcomeType = -1
)

type SpecifierType int8

const (
	SpecifierTypeString SpecifierType = iota
	SpecifierTypeInteger
	SpecifierTypeDecimal
	SpecifierTypeVariableText
	SpecifierTypeUnknown SpecifierType = -1
)

type Gender int8

const (
	GenderUnknown Gender = iota
	Male
	Female
)
