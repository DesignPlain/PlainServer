package iam

type GroupPolicyAttachment struct {
	// The group the policy should be applied to
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// The ARN of the policy you want to apply
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`
}
