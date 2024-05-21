package types

type Compute_getInstanceTemplateReservationAffinitySpecificReservation struct {
	// The key for the node affinity label.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Corresponds to the label values of a reservation resource.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
