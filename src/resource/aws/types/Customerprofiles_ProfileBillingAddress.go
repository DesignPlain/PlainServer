package types

type Customerprofiles_ProfileBillingAddress struct {
	// The first line of a customer address.
	Address1 string `json:"address1,omitempty" yaml:"address1,omitempty"`

	// The third line of a customer address.
	Address3 string `json:"address3,omitempty" yaml:"address3,omitempty"`

	// The fourth line of a customer address.
	Address4 string `json:"address4,omitempty" yaml:"address4,omitempty"`

	// The city in which a customer lives.
	City string `json:"city,omitempty" yaml:"city,omitempty"`

	// The postal code of a customer address.
	PostalCode string `json:"postalCode,omitempty" yaml:"postalCode,omitempty"`

	// The province in which a customer lives.
	Province string `json:"province,omitempty" yaml:"province,omitempty"`

	// The state in which a customer lives.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// The second line of a customer address.
	Address2 string `json:"address2,omitempty" yaml:"address2,omitempty"`

	// The country in which a customer lives.
	Country string `json:"country,omitempty" yaml:"country,omitempty"`

	// The county in which a customer lives.
	County string `json:"county,omitempty" yaml:"county,omitempty"`
}
