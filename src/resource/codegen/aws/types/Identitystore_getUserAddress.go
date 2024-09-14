package types

type Identitystore_getUserAddress struct {
	// The region of the address.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The street of the address.
	StreetAddress string `json:"streetAddress,omitempty" yaml:"streetAddress,omitempty"`

	// The type of phone number.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The country that this address is in.
	Country string `json:"country,omitempty" yaml:"country,omitempty"`

	// The name that is typically displayed when the name is shown for display.
	Formatted string `json:"formatted,omitempty" yaml:"formatted,omitempty"`

	// The address locality.
	Locality string `json:"locality,omitempty" yaml:"locality,omitempty"`

	// The postal code of the address.
	PostalCode string `json:"postalCode,omitempty" yaml:"postalCode,omitempty"`

	// When `true`, this is the primary phone number associated with the user.
	Primary bool `json:"primary,omitempty" yaml:"primary,omitempty"`
}
