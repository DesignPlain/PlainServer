package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity struct {
	// Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION
	ConsumeReservationType string `json:"consumeReservationType,omitempty" yaml:"consumeReservationType,omitempty"`

	// Corresponds to the label key of reservation resource.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Corresponds to the label values of reservation resource.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
