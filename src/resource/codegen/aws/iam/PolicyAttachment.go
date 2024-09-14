package iam

type PolicyAttachment struct {
	// Group(s) the policy should be applied to.
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	// Name of the attachment. This cannot be an empty string.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ARN of the policy you want to apply. Typically this should be a reference to the ARN of another resource to ensure dependency ordering, such as `aws_iam_policy.example.arn`.
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`

	// Role(s) the policy should be applied to.
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`

	// User(s) the policy should be applied to.
	Users []string `json:"users,omitempty" yaml:"users,omitempty"`
}
