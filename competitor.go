package uof

import (
	"time"
)

type CompetitorProfile struct {
	Competitor  CompetitorPlayer `xml:"competitor" json:"competitor" bson:"competitor,omitempty"`
	GeneratedAt time.Time        `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty" bson:"generatedAt,omitempty"`
}
