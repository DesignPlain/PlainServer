package organizations

import types "DesignSphere_Server/src/resource/gcp/types"

type IAMBinding struct {
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Organizations_IAMBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
