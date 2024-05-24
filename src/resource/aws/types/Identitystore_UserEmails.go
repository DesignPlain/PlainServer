package types

type Identitystore_UserEmails struct {
	// The type of email.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The email address. This value must be unique across the identity store.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// When `true`, this is the primary email associated with the user.
	Primary bool `json:"primary,omitempty" yaml:"primary,omitempty"`
}
