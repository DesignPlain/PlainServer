package types

type Lambda_EventSourceMappingDestinationConfig struct {
	// The destination configuration for failed invocations. Detailed below.
	OnFailure Lambda_EventSourceMappingDestinationConfigOnFailure `json:"onFailure,omitempty" yaml:"onFailure,omitempty"`
}
