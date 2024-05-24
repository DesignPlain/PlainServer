package types

type Index_ProviderAssumeRoleWithWebIdentity struct {
	//
	WebIdentityTokenFile string `json:"webIdentityTokenFile,omitempty" yaml:"webIdentityTokenFile,omitempty"`

	// The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	PolicyArns []string `json:"policyArns,omitempty" yaml:"policyArns,omitempty"`

	// Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// An identifier for the assumed role session.
	SessionName string `json:"sessionName,omitempty" yaml:"sessionName,omitempty"`

	//
	WebIdentityToken string `json:"webIdentityToken,omitempty" yaml:"webIdentityToken,omitempty"`
}
