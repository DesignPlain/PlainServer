package types

type Elasticache_UserAuthenticationMode struct {
	// Specifies the passwords to use for authentication if `type` is set to `password`.
	Passwords []string `json:"passwords,omitempty" yaml:"passwords,omitempty"`

	// Specifies the authentication type. Possible options are: `password`, `no-password-required` or `iam`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	PasswordCount int `json:"passwordCount,omitempty" yaml:"passwordCount,omitempty"`
}
