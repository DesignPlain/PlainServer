package acmpca

type Policy struct {
	// JSON-formatted IAM policy to attach to the specified private CA resource.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// ARN of the private CA to associate with the policy.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
