package types

type Compute_InstanceReservationAffinity struct {
	/*
	   Specifies the label selector for the reservation to use..
	   Structure is documented below.
	*/
	SpecificReservation Compute_InstanceReservationAffinitySpecificReservation `json:"specificReservation,omitempty" yaml:"specificReservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
