package sesv2

type EmailIdentityPolicy struct {
	// The email identity.
	EmailIdentity string `json:"emailIdentity,omitempty" yaml:"emailIdentity,omitempty"`

	// The text of the policy in JSON format.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The name of the policy.
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}
