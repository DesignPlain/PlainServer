package types

type Identityplatform_ConfigSignIn struct {
	// Whether to allow more than one account to have the same email.
	AllowDuplicateEmails bool `json:"allowDuplicateEmails,omitempty" yaml:"allowDuplicateEmails,omitempty"`

	/*
	   Configuration options related to authenticating an anonymous user.
	   Structure is documented below.
	*/
	Anonymous Identityplatform_ConfigSignInAnonymous `json:"anonymous,omitempty" yaml:"anonymous,omitempty"`

	/*
	   Configuration options related to authenticating a user by their email address.
	   Structure is documented below.
	*/
	Email Identityplatform_ConfigSignInEmail `json:"email,omitempty" yaml:"email,omitempty"`

	/*
	   (Output)
	   Output only. Hash config information.
	   Structure is documented below.
	*/
	HashConfigs []Identityplatform_ConfigSignInHashConfig `json:"hashConfigs,omitempty" yaml:"hashConfigs,omitempty"`

	/*
	   Configuration options related to authenticated a user by their phone number.
	   Structure is documented below.
	*/
	PhoneNumber Identityplatform_ConfigSignInPhoneNumber `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
}
