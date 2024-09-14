package organizations

import types "libds/gcp/types"

type IAMMember struct {
	/*
	   An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	   Structure is documented below.
	*/
	Condition types.Organizations_IAMMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	// The organization id of the target organization.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	   `organizations/{{org_id}}/roles/{{role_id}}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
