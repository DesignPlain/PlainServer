package iam

type AccountPasswordPolicy struct {
	// Whether to allow users to change their own password
	AllowUsersToChangePassword bool `json:"allowUsersToChangePassword,omitempty" yaml:"allowUsersToChangePassword,omitempty"`

	// The number of days that an user password is valid.
	MaxPasswordAge int `json:"maxPasswordAge,omitempty" yaml:"maxPasswordAge,omitempty"`

	// Whether to require symbols for user passwords.
	RequireSymbols bool `json:"requireSymbols,omitempty" yaml:"requireSymbols,omitempty"`

	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry bool `json:"hardExpiry,omitempty" yaml:"hardExpiry,omitempty"`

	// Minimum length to require for user passwords.
	MinimumPasswordLength int `json:"minimumPasswordLength,omitempty" yaml:"minimumPasswordLength,omitempty"`

	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention int `json:"passwordReusePrevention,omitempty" yaml:"passwordReusePrevention,omitempty"`

	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters bool `json:"requireLowercaseCharacters,omitempty" yaml:"requireLowercaseCharacters,omitempty"`

	// Whether to require numbers for user passwords.
	RequireNumbers bool `json:"requireNumbers,omitempty" yaml:"requireNumbers,omitempty"`

	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters bool `json:"requireUppercaseCharacters,omitempty" yaml:"requireUppercaseCharacters,omitempty"`
}
