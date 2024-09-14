package appmesh

import types "libds/aws/types"

type Mesh struct {
	// Name to use for the service mesh. Must be between 1 and 255 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Service mesh specification to apply.
	Spec types.Appmesh_MeshSpec `json:"spec,omitempty" yaml:"spec,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
