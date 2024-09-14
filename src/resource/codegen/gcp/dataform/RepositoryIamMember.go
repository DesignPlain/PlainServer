package dataform

import types "libds/gcp/types"

type RepositoryIamMember struct {
	//
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Dataform_RepositoryIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
