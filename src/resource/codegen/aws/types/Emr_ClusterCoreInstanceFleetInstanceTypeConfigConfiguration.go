package types

type Emr_ClusterCoreInstanceFleetInstanceTypeConfigConfiguration struct {
	// Map of properties specified within a configuration classification.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Classification within a configuration.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`
}
