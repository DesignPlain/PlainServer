package neptune

import types "DesignSphere_Server/src/resource/aws/types"

type ParameterGroup struct {
	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the Neptune parameter group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The family of the Neptune parameter group.
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	// The name of the Neptune parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// A list of Neptune parameters to apply.
	Parameters []types.Neptune_ParameterGroupParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
