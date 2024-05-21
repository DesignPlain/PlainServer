package types

type Clouddomains_RegistrationContactSettings struct {
	/*
	   Required. Privacy setting for the contacts associated with the Registration.
	   Values are PUBLIC_CONTACT_DATA, PRIVATE_CONTACT_DATA, and REDACTED_CONTACT_DATA
	*/
	Privacy string `json:"privacy,omitempty" yaml:"privacy,omitempty"`

	/*
	   Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.
	   Warning: For new Registrations, the registrant receives an email confirmation that they must complete within 15 days to
	   avoid domain suspension.
	   Structure is documented below.
	*/
	RegistrantContact Clouddomains_RegistrationContactSettingsRegistrantContact `json:"registrantContact,omitempty" yaml:"registrantContact,omitempty"`

	/*
	   Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.
	   Warning: For new Registrations, the registrant receives an email confirmation that they must complete within 15 days to
	   avoid domain suspension.
	   Structure is documented below.
	*/
	TechnicalContact Clouddomains_RegistrationContactSettingsTechnicalContact `json:"technicalContact,omitempty" yaml:"technicalContact,omitempty"`

	/*
	   Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.
	   Warning: For new Registrations, the registrant receives an email confirmation that they must complete within 15 days to
	   avoid domain suspension.
	   Structure is documented below.
	*/
	AdminContact Clouddomains_RegistrationContactSettingsAdminContact `json:"adminContact,omitempty" yaml:"adminContact,omitempty"`
}
