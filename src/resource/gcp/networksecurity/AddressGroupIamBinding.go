package networksecurity

import types "DesignSphere_Server/src/resource/gcp/types"

type AddressGroupIamBinding struct {
	//
	Condition types.Networksecurity_AddressGroupIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
