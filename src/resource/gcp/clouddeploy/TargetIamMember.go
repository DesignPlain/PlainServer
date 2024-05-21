package clouddeploy

import types "DesignSphere_Server/src/resource/gcp/types"

type TargetIamMember struct {
	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Clouddeploy_TargetIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
