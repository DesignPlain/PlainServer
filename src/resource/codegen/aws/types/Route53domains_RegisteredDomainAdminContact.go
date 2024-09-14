package types

type Route53domains_RegisteredDomainAdminContact struct {
	// Indicates whether the contact is a person, company, association, or public organization. See the [AWS API documentation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ContactDetail.html#Route53Domains-Type-domains_ContactDetail-ContactType) for valid values.
	ContactType string `json:"contactType,omitempty" yaml:"contactType,omitempty"`

	// A key-value map of parameters required by certain top-level domains.
	ExtraParams map[string]string `json:"extraParams,omitempty" yaml:"extraParams,omitempty"`

	// Last name of contact.
	LastName string `json:"lastName,omitempty" yaml:"lastName,omitempty"`

	// First line of the contact's address.
	AddressLine1 string `json:"addressLine1,omitempty" yaml:"addressLine1,omitempty"`

	// The city of the contact's address.
	City string `json:"city,omitempty" yaml:"city,omitempty"`

	// Code for the country of the contact's address. See the [AWS API documentation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ContactDetail.html#Route53Domains-Type-domains_ContactDetail-CountryCode) for valid values.
	CountryCode string `json:"countryCode,omitempty" yaml:"countryCode,omitempty"`

	// First name of contact.
	FirstName string `json:"firstName,omitempty" yaml:"firstName,omitempty"`

	// Name of the organization for contact types other than `PERSON`.
	OrganizationName string `json:"organizationName,omitempty" yaml:"organizationName,omitempty"`

	// The state or province of the contact's city.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// The zip or postal code of the contact's address.
	ZipCode string `json:"zipCode,omitempty" yaml:"zipCode,omitempty"`

	// Second line of contact's address, if any.
	AddressLine2 string `json:"addressLine2,omitempty" yaml:"addressLine2,omitempty"`

	// Email address of the contact.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Fax number of the contact. Phone number must be specified in the format "+[country dialing code].[number including any area code]".
	Fax string `json:"fax,omitempty" yaml:"fax,omitempty"`

	// The phone number of the contact. Phone number must be specified in the format "+[country dialing code].[number including any area code]".
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
}
