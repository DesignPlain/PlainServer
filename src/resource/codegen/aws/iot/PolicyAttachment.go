package iot

type PolicyAttachment struct {
	// The name of the policy to attach.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The identity to which the policy is attached.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
