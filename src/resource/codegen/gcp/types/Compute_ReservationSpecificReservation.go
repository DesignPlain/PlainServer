package types

type Compute_ReservationSpecificReservation struct {
	// The number of resources that are allocated.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	/*
	   (Output)
	   How many instances are in use.
	*/
	InUseCount int `json:"inUseCount,omitempty" yaml:"inUseCount,omitempty"`

	/*
	   The instance properties for the reservation.
	   Structure is documented below.
	*/
	InstanceProperties Compute_ReservationSpecificReservationInstanceProperties `json:"instanceProperties,omitempty" yaml:"instanceProperties,omitempty"`
}
