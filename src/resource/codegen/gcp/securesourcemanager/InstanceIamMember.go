package securesourcemanager

import types "libds/gcp/types"

type InstanceIamMember struct {
	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

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
}
