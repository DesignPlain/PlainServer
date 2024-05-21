package securitycenter

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceIamMember struct {
	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	// The ID of the instance or a fully qualified identifier for the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the Data Fusion instance.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Securitycenter_InstanceIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
