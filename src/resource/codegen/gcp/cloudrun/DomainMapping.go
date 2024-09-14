package cloudrun

import types "libds/gcp/types"

type DomainMapping struct {
	// The location of the cloud run instance. eg us-central1
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Metadata associated with this DomainMapping.
	   Structure is documented below.
	*/
	Metadata types.Cloudrun_DomainMappingMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The spec for this DomainMapping.
	   Structure is documented below.
	*/
	Spec types.Cloudrun_DomainMappingSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}
