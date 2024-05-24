package account

type PrimaryContact struct {
	// The full name of the primary contact address.
	FullName string `json:"fullName,omitempty" yaml:"fullName,omitempty"`

	// The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`

	// The state or region of the primary contact address. This field is required in selected countries.
	StateOrRegion string `json:"stateOrRegion,omitempty" yaml:"stateOrRegion,omitempty"`

	// The second line of the primary contact address, if any.
	AddressLine2 string `json:"addressLine2,omitempty" yaml:"addressLine2,omitempty"`

	// The ISO-3166 two-letter country code for the primary contact address.
	CountryCode string `json:"countryCode,omitempty" yaml:"countryCode,omitempty"`

	// The district or county of the primary contact address, if any.
	DistrictOrCounty string `json:"districtOrCounty,omitempty" yaml:"districtOrCounty,omitempty"`

	// The city of the primary contact address.
	City string `json:"city,omitempty" yaml:"city,omitempty"`

	// The name of the company associated with the primary contact information, if any.
	CompanyName string `json:"companyName,omitempty" yaml:"companyName,omitempty"`

	// The postal code of the primary contact address.
	PostalCode string `json:"postalCode,omitempty" yaml:"postalCode,omitempty"`

	// The URL of the website associated with the primary contact information, if any.
	WebsiteUrl string `json:"websiteUrl,omitempty" yaml:"websiteUrl,omitempty"`

	// The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The first line of the primary contact address.
	AddressLine1 string `json:"addressLine1,omitempty" yaml:"addressLine1,omitempty"`

	// The third line of the primary contact address, if any.
	AddressLine3 string `json:"addressLine3,omitempty" yaml:"addressLine3,omitempty"`
}
