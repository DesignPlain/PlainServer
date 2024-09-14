package ses

type IdentityPolicy struct {
	// Name or Amazon Resource Name (ARN) of the SES Identity.
	Identity string `json:"identity,omitempty" yaml:"identity,omitempty"`

	// Name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// JSON string of the policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
