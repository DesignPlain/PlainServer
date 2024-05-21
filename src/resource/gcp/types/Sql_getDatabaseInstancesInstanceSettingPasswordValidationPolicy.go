package types

type Sql_getDatabaseInstancesInstanceSettingPasswordValidationPolicy struct {
	// Number of previous passwords that cannot be reused.
	ReuseInterval int `json:"reuseInterval,omitempty" yaml:"reuseInterval,omitempty"`

	// Password complexity.
	Complexity string `json:"complexity,omitempty" yaml:"complexity,omitempty"`

	// Disallow username as a part of the password.
	DisallowUsernameSubstring bool `json:"disallowUsernameSubstring,omitempty" yaml:"disallowUsernameSubstring,omitempty"`

	// Whether the password policy is enabled or not.
	EnablePasswordPolicy bool `json:"enablePasswordPolicy,omitempty" yaml:"enablePasswordPolicy,omitempty"`

	// Minimum number of characters allowed.
	MinLength int `json:"minLength,omitempty" yaml:"minLength,omitempty"`

	// Minimum interval after which the password can be changed. This flag is only supported for PostgresSQL.
	PasswordChangeInterval string `json:"passwordChangeInterval,omitempty" yaml:"passwordChangeInterval,omitempty"`
}
