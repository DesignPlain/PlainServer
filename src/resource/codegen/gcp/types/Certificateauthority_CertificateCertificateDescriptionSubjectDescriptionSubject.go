package types

type Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubject struct {
	// The locality or city of the subject.
	Locality string `json:"locality,omitempty" yaml:"locality,omitempty"`

	// The organization of the subject.
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	// The organizational unit of the subject.
	OrganizationalUnit string `json:"organizationalUnit,omitempty" yaml:"organizationalUnit,omitempty"`

	// The postal code of the subject.
	PostalCode string `json:"postalCode,omitempty" yaml:"postalCode,omitempty"`

	// The province, territory, or regional state of the subject.
	Province string `json:"province,omitempty" yaml:"province,omitempty"`

	// The street address of the subject.
	StreetAddress string `json:"streetAddress,omitempty" yaml:"streetAddress,omitempty"`

	// The common name of the distinguished name.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	// The country code of the subject.
	CountryCode string `json:"countryCode,omitempty" yaml:"countryCode,omitempty"`
}
