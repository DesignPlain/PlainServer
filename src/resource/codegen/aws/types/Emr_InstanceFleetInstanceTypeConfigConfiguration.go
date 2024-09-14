package types

type Emr_InstanceFleetInstanceTypeConfigConfiguration struct {
	// The classification within a configuration.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// A map of properties specified within a configuration classification
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
