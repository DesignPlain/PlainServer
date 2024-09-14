package securesourcemanager

import types "libds/gcp/types"

type InstanceIamBinding struct {
	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Securesourcemanager_InstanceIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
