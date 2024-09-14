package types

type Clouddomains_RegistrationContactSettingsRegistrantContact struct {
	/*
	   Required. Postal address of the contact.
	   Structure is documented below.
	*/
	PostalAddress Clouddomains_RegistrationContactSettingsRegistrantContactPostalAddress `json:"postalAddress,omitempty" yaml:"postalAddress,omitempty"`

	// Required. Email address of the contact.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Fax number of the contact in international format. For example, "+1-800-555-0123".
	FaxNumber string `json:"faxNumber,omitempty" yaml:"faxNumber,omitempty"`

	// Required. Phone number of the contact in international format. For example, "+1-800-555-0123".
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
}
