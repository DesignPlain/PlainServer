package dataform

import types "libds/gcp/types"

type RepositoryIamBinding struct {
	//
	Condition types.Dataform_RepositoryIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
