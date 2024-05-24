package types

type Emr_ClusterMasterInstanceFleetInstanceTypeConfigConfiguration struct {
	// Classification within a configuration.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// Map of properties specified within a configuration classification.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
