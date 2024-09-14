package types

type Cloudrun_ServiceTemplateSpecContainerEnvFromSecretRef struct {
	// Specify whether the Secret must be defined
	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`

	/*
	   The Secret to select from.
	   Structure is documented below.
	*/
	LocalObjectReference Cloudrun_ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference `json:"localObjectReference,omitempty" yaml:"localObjectReference,omitempty"`
}
