package securesourcemanager

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceIamMember struct {
	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Securesourcemanager_InstanceIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
}
