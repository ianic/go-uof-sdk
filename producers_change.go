package uof

type ProducersChange []ProducerChange

type ProducerChange struct {
	Producer   Producer       `json:"producer" bson:"producer"`
	Status     ProducerStatus `json:"status" bson:"status"`
	RecoveryID int            `json:"recoveryId,omitempty" bson:"recoveryId,omitempty"`
	Timestamp  int            `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

func (p *ProducersChange) Add(producer Producer, timestamp int) {
	*p = append(*p, ProducerChange{Producer: producer, Timestamp: timestamp})
}

func (p *ProducersChange) AddAll(producers []Producer, timestamp int) {
	for _, d := range producers {
		p.Add(d, timestamp)
	}
}

type ProducerStatus int8

const (
	ProducerStatusDown       ProducerStatus = -1
	ProducerStatusActive     ProducerStatus = 1
	ProducerStatusInRecovery ProducerStatus = 2
)
