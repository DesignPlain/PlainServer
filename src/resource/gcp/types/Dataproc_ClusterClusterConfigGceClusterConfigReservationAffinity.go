package types

type Dataproc_ClusterClusterConfigGceClusterConfigReservationAffinity struct {
	// Corresponds to the type of reservation consumption.
	ConsumeReservationType string `json:"consumeReservationType,omitempty" yaml:"consumeReservationType,omitempty"`

	// Corresponds to the label key of reservation resource.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Corresponds to the label values of reservation resource.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
