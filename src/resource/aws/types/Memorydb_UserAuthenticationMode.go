package types

type Memorydb_UserAuthenticationMode struct {
	// Indicates whether the user requires a password to authenticate. Must be set to `password`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The number of passwords belonging to the user.
	PasswordCount int `json:"passwordCount,omitempty" yaml:"passwordCount,omitempty"`

	// The set of passwords used for authentication. You can create up to two passwords for each user.
	Passwords []string `json:"passwords,omitempty" yaml:"passwords,omitempty"`
}
