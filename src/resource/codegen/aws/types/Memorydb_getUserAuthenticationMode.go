package types

type Memorydb_getUserAuthenticationMode struct {
	// Number of passwords belonging to the user if `type` is set to `password`.
	PasswordCount int `json:"passwordCount,omitempty" yaml:"passwordCount,omitempty"`

	// Type of authentication configured.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
