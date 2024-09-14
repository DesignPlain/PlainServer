package types

type Compute_RegionInstanceTemplateReservationAffinity struct {
	/*
	   Specifies the label selector for the reservation to use..
	   Structure is documented below.
	*/
	SpecificReservation Compute_RegionInstanceTemplateReservationAffinitySpecificReservation `json:"specificReservation,omitempty" yaml:"specificReservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
