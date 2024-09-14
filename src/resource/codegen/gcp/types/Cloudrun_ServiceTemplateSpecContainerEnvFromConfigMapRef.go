package types

type Cloudrun_ServiceTemplateSpecContainerEnvFromConfigMapRef struct {
	/*
	   The ConfigMap to select from.
	   Structure is documented below.
	*/
	LocalObjectReference Cloudrun_ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference `json:"localObjectReference,omitempty" yaml:"localObjectReference,omitempty"`

	// Specify whether the ConfigMap must be defined
	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`
}
