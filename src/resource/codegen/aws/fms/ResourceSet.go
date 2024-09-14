package fms

import types "libds/aws/types"

type ResourceSet struct {
	// Details about the resource set to be created or updated. See `resource_set` Attribute Reference below.
	ResourceSets []types.Fms_ResourceSetResourceSet `json:"resourceSets,omitempty" yaml:"resourceSets,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Fms_ResourceSetTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
