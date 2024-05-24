package types

type Identitystore_getUserEmail struct {
	// The user's phone number.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// When `true`, this is the primary phone number associated with the user.
	Primary bool `json:"primary,omitempty" yaml:"primary,omitempty"`

	// The type of phone number.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
