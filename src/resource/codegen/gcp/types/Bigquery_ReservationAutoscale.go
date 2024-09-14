package types

type Bigquery_ReservationAutoscale struct {
	/*
	   (Output)
	   The slot capacity added to this reservation when autoscale happens. Will be between [0, max_slots].
	*/
	CurrentSlots int `json:"currentSlots,omitempty" yaml:"currentSlots,omitempty"`

	// Number of slots to be scaled when needed.
	MaxSlots int `json:"maxSlots,omitempty" yaml:"maxSlots,omitempty"`
}
