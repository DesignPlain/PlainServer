package types

type Mediaconvert_QueueReservationPlanSettings struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment string `json:"commitment,omitempty" yaml:"commitment,omitempty"`

	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType string `json:"renewalType,omitempty" yaml:"renewalType,omitempty"`

	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots int `json:"reservedSlots,omitempty" yaml:"reservedSlots,omitempty"`
}
