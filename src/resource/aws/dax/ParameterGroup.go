package dax

import types "DesignSphere_Server/src/resource/aws/types"

type ParameterGroup struct {
	// A description of the parameter group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the parameter group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parameters of the parameter group.
	Parameters []types.Dax_ParameterGroupParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
