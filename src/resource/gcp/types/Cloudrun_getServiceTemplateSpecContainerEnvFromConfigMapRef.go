package types



type Cloudrun_getServiceTemplateSpecContainerEnvFromConfigMapRef struct {
	// The ConfigMap to select from.
	LocalObjectReferences []Cloudrun_getServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference `json:"localObjectReferences,omitempty" yaml:"localObjectReferences,omitempty"`

	// Specify whether the ConfigMap must be defined
	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`
}
