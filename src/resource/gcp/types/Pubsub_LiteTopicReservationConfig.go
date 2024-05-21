package types

type Pubsub_LiteTopicReservationConfig struct {
	// The Reservation to use for this topic's throughput capacity.
	ThroughputReservation string `json:"throughputReservation,omitempty" yaml:"throughputReservation,omitempty"`
}
