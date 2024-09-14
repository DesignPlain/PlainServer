package projects

import types "libds/gcp/types"

type IAMBinding struct {
	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The project id of the target project. This is not
	   inferred from the provider.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	/*
	   An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	   Structure is documented below.
	*/
	Condition types.Projects_IAMBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
