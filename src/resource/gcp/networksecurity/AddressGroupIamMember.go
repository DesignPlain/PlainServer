package networksecurity

import types "DesignSphere_Server/src/resource/gcp/types"

type AddressGroupIamMember struct {
	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Networksecurity_AddressGroupIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
