package vpclattice

type ResourcePolicy struct {
	// An IAM policy. The policy string in JSON must not contain newlines or blank lines.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The ID or Amazon Resource Name (ARN) of the service network or service for which the policy is created.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
