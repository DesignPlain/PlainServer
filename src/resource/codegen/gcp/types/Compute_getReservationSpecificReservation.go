package types

type Compute_getReservationSpecificReservation struct {
	// The number of resources that are allocated.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// How many instances are in use.
	InUseCount int `json:"inUseCount,omitempty" yaml:"inUseCount,omitempty"`

	// The instance properties for the reservation.
	InstanceProperties []Compute_getReservationSpecificReservationInstanceProperty `json:"instanceProperties,omitempty" yaml:"instanceProperties,omitempty"`
}
