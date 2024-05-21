package types

type Cloudrun_getServiceTemplateSpecContainerEnvFromSecretRef struct {
	// The Secret to select from.
	LocalObjectReferences []Cloudrun_getServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference `json:"localObjectReferences,omitempty" yaml:"localObjectReferences,omitempty"`

	// Specify whether the Secret must be defined
	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`
}
