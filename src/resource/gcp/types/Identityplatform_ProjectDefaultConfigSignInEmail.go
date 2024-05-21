package types

type Identityplatform_ProjectDefaultConfigSignInEmail struct {
	// Whether email auth is enabled for the project or not.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   Whether a password is required for email auth or not. If true, both an email and
	   password must be provided to sign in. If false, a user may sign in via either
	   email/password or email link.
	*/
	PasswordRequired bool `json:"passwordRequired,omitempty" yaml:"passwordRequired,omitempty"`
}
