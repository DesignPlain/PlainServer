package organizations

type PolicyAttachment struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`

	// If set to `true`, destroy will --not-- detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`
}
