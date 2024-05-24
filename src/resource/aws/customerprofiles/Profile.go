package customerprofiles

import types "DesignSphere_Server/src/resource/aws/types"

type Profile struct {
	// A unique account number that you have given to the customer.
	AccountNumber string `json:"accountNumber,omitempty" yaml:"accountNumber,omitempty"`

	// A block that specifies the customer’s billing address. Documented below.
	BillingAddress types.Customerprofiles_ProfileBillingAddress `json:"billingAddress,omitempty" yaml:"billingAddress,omitempty"`

	// The customer’s business email address.
	BusinessEmailAddress string `json:"businessEmailAddress,omitempty" yaml:"businessEmailAddress,omitempty"`

	/*
	   The name of your Customer Profile domain. It must be unique for your AWS account.

	   The following arguments are optional:
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The customer’s mobile phone number.
	MobilePhoneNumber string `json:"mobilePhoneNumber,omitempty" yaml:"mobilePhoneNumber,omitempty"`

	// The type of profile used to describe the customer.
	PartyTypeString string `json:"partyTypeString,omitempty" yaml:"partyTypeString,omitempty"`

	// Any additional information relevant to the customer’s profile.
	AdditionalInformation string `json:"additionalInformation,omitempty" yaml:"additionalInformation,omitempty"`

	// The customer’s birth date.
	BirthDate string `json:"birthDate,omitempty" yaml:"birthDate,omitempty"`

	// The customer’s business phone number.
	BusinessPhoneNumber string `json:"businessPhoneNumber,omitempty" yaml:"businessPhoneNumber,omitempty"`

	// The customer’s email address, which has not been specified as a personal or business address.
	EmailAddress string `json:"emailAddress,omitempty" yaml:"emailAddress,omitempty"`

	// The gender with which the customer identifies.
	GenderString string `json:"genderString,omitempty" yaml:"genderString,omitempty"`

	// The customer’s last name.
	LastName string `json:"lastName,omitempty" yaml:"lastName,omitempty"`

	// The customer’s personal email address.
	PersonalEmailAddress string `json:"personalEmailAddress,omitempty" yaml:"personalEmailAddress,omitempty"`

	// The customer’s phone number, which has not been specified as a mobile, home, or business number.
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`

	// The customer’s middle name.
	MiddleName string `json:"middleName,omitempty" yaml:"middleName,omitempty"`

	// A block that specifies the customer’s shipping address. Documented below.
	ShippingAddress types.Customerprofiles_ProfileShippingAddress `json:"shippingAddress,omitempty" yaml:"shippingAddress,omitempty"`

	// A block that specifies a generic address associated with the customer that is not mailing, shipping, or billing. Documented below.
	Address types.Customerprofiles_ProfileAddress `json:"address,omitempty" yaml:"address,omitempty"`

	// A key value pair of attributes of a customer profile.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// The name of the customer’s business.
	BusinessName string `json:"businessName,omitempty" yaml:"businessName,omitempty"`

	// The customer’s first name.
	FirstName string `json:"firstName,omitempty" yaml:"firstName,omitempty"`

	// The customer’s home phone number.
	HomePhoneNumber string `json:"homePhoneNumber,omitempty" yaml:"homePhoneNumber,omitempty"`

	// A block that specifies the customer’s mailing address. Documented below.
	MailingAddress types.Customerprofiles_ProfileMailingAddress `json:"mailingAddress,omitempty" yaml:"mailingAddress,omitempty"`
}
