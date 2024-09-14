package vpclattice

type AuthPolicy struct {
	// The state of the auth policy. The auth policy is only active when the auth type is set to `AWS_IAM`. If you provide a policy, then authentication and authorization decisions are made based on this policy and the client's IAM policy. If the Auth type is `NONE`, then, any auth policy you provide will remain inactive.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// The auth policy. The policy string in JSON must not contain newlines or blank lines.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The ID or Amazon Resource Name (ARN) of the service network or service for which the policy is created.
	ResourceIdentifier string `json:"resourceIdentifier,omitempty" yaml:"resourceIdentifier,omitempty"`
}
