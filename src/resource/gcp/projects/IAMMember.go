package projects

import types "DesignSphere_Server/src/resource/gcp/types"

type IAMMember struct {
	/*
	   An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	   Structure is documented below.
	*/
	Condition types.Projects_IAMMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

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
}
