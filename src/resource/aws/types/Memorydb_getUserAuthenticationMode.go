package types

type Memorydb_getUserAuthenticationMode struct {
	// The number of passwords belonging to the user.
	PasswordCount int `json:"passwordCount,omitempty" yaml:"passwordCount,omitempty"`

	// Whether the user requires a password to authenticate.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
