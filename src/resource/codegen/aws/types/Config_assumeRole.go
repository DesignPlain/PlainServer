package types

type Config_assumeRole struct {
	// Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// An identifier for the assumed role session.
	SessionName string `json:"sessionName,omitempty" yaml:"sessionName,omitempty"`

	// Source identity specified by the principal assuming the role.
	SourceIdentity string `json:"sourceIdentity,omitempty" yaml:"sourceIdentity,omitempty"`

	// Assume role session tags.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Assume role session tag keys to pass to any subsequent sessions.
	TransitiveTagKeys []string `json:"transitiveTagKeys,omitempty" yaml:"transitiveTagKeys,omitempty"`

	// A unique identifier that might be required when you assume a role in another account.
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	PolicyArns []string `json:"policyArns,omitempty" yaml:"policyArns,omitempty"`
}
