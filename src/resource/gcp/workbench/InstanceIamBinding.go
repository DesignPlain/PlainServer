package workbench

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceIamBinding struct {
	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Workbench_InstanceIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`
}
