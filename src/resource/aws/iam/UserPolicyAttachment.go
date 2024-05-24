package iam

type UserPolicyAttachment struct {
	// The ARN of the policy you want to apply
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`

	// The user the policy should be applied to
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
