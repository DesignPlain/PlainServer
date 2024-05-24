package iot

type RoleAlias struct {
	// The identity of the role to which the alias refers.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The name of the role alias.
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	CredentialDuration int `json:"credentialDuration,omitempty" yaml:"credentialDuration,omitempty"`
}
