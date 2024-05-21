package securesourcemanager

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceIamBinding struct {
	//
	Condition types.Securesourcemanager_InstanceIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
