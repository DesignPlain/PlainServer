package types

type Compute_getInstanceReservationAffinity struct {
	// Specifies the label selector for the reservation to use.
	SpecificReservations []Compute_getInstanceReservationAffinitySpecificReservation `json:"specificReservations,omitempty" yaml:"specificReservations,omitempty"`

	// The accelerator type resource exposed to this instance. E.g. `nvidia-tesla-k80`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
