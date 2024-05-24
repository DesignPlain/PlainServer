package iam

type PolicyAttachment struct {
	// The group(s) the policy should be applied to
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	// The name of the attachment. This cannot be an empty string.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of the policy you want to apply
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`

	// The role(s) the policy should be applied to
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`

	// The user(s) the policy should be applied to
	Users []string `json:"users,omitempty" yaml:"users,omitempty"`
}
