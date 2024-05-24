package types

type Connect_getUserIdentityInfo struct {
	// The email address.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// The first name.
	FirstName string `json:"firstName,omitempty" yaml:"firstName,omitempty"`

	// The last name.
	LastName string `json:"lastName,omitempty" yaml:"lastName,omitempty"`
}
