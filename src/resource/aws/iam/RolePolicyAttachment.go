package iam

type RolePolicyAttachment struct {
	// The ARN of the policy you want to apply
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`

	// The name of the IAM role to which the policy should be applied
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
