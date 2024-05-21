package types

type Compute_InstanceFromMachineImageReservationAffinity struct {
	// Specifies the label selector for the reservation to use.
	SpecificReservation Compute_InstanceFromMachineImageReservationAffinitySpecificReservation `json:"specificReservation,omitempty" yaml:"specificReservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
