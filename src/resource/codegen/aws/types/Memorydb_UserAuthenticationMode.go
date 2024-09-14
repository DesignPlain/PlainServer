package types

type Memorydb_UserAuthenticationMode struct {
	// Number of passwords belonging to the user if `type` is set to `password`.
	PasswordCount int `json:"passwordCount,omitempty" yaml:"passwordCount,omitempty"`

	// Set of passwords used for authentication if `type` is set to `password`. You can create up to two passwords for each user.
	Passwords []string `json:"passwords,omitempty" yaml:"passwords,omitempty"`

	// Specifies the authentication type. Valid values are: `password` or `iam`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
