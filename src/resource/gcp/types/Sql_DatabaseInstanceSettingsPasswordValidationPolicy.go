package types

type Sql_DatabaseInstanceSettingsPasswordValidationPolicy struct {
	/*
	   Enables or disable the password validation policy.

	   The optional `replica_configuration` block must have `master_instance_name` set
	   to work, cannot be updated, and supports:
	*/
	EnablePasswordPolicy bool `json:"enablePasswordPolicy,omitempty" yaml:"enablePasswordPolicy,omitempty"`

	// Specifies the minimum number of characters that the password must have.
	MinLength int `json:"minLength,omitempty" yaml:"minLength,omitempty"`

	// Specifies the minimum duration after which you can change the password.
	PasswordChangeInterval string `json:"passwordChangeInterval,omitempty" yaml:"passwordChangeInterval,omitempty"`

	// Specifies the number of previous passwords that you can't reuse.
	ReuseInterval int `json:"reuseInterval,omitempty" yaml:"reuseInterval,omitempty"`

	// Checks if the password is a combination of lowercase, uppercase, numeric, and non-alphanumeric characters.
	Complexity string `json:"complexity,omitempty" yaml:"complexity,omitempty"`

	// Prevents the use of the username in the password.
	DisallowUsernameSubstring bool `json:"disallowUsernameSubstring,omitempty" yaml:"disallowUsernameSubstring,omitempty"`
}
