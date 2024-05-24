package iam

type UserLoginProfile struct {
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
	PasswordLength int `json:"passwordLength,omitempty" yaml:"passwordLength,omitempty"`

	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	PasswordResetRequired bool `json:"passwordResetRequired,omitempty" yaml:"passwordResetRequired,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey string `json:"pgpKey,omitempty" yaml:"pgpKey,omitempty"`

	// The IAM user's name.
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
