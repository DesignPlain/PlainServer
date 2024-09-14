package oam

type SinkPolicy struct {
	// JSON policy to use. If you are updating an existing policy, the entire existing policy is replaced by what you specify here.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// ARN of the sink to attach this policy to.
	SinkIdentifier string `json:"sinkIdentifier,omitempty" yaml:"sinkIdentifier,omitempty"`
}
