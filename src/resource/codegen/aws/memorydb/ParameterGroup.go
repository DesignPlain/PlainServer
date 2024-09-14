package memorydb

import types "libds/aws/types"

type ParameterGroup struct {
	// Description for the parameter group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The engine version that the parameter group can be used with.

	   The following arguments are optional:
	*/
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	// Name of the parameter group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
	Parameters []types.Memorydb_ParameterGroupParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
