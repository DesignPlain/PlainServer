package types

type Apigee_KeystoresAliasesSelfSignedCertSubject struct {
	// Organization name. Maximum length is 64 characters.
	Org string `json:"org,omitempty" yaml:"org,omitempty"`

	// Organization team name. Maximum length is 64 characters.
	OrgUnit string `json:"orgUnit,omitempty" yaml:"orgUnit,omitempty"`

	// State or district name. Maximum length is 128 characters.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Common name of the organization. Maximum length is 64 characters.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	// Two-letter country code. Example, IN for India, US for United States of America.
	CountryCode string `json:"countryCode,omitempty" yaml:"countryCode,omitempty"`

	/*
	   Email address. Max 255 characters.

	   - - -
	*/
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// City or town name. Maximum length is 128 characters.
	Locality string `json:"locality,omitempty" yaml:"locality,omitempty"`
}
