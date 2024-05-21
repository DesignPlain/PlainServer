package types

type Container_getClusterNodePoolNodeConfigReservationAffinity struct {
	// Corresponds to the type of reservation consumption.
	ConsumeReservationType string `json:"consumeReservationType,omitempty" yaml:"consumeReservationType,omitempty"`

	// The label key of a reservation resource.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The label values of the reservation resource.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
