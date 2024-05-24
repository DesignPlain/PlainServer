package types

type Acmpca_CertificateAuthorityCertificateAuthorityConfigurationSubject struct {
	// Fully qualified domain name (FQDN) associated with the certificate subject. Must be less than or equal to 64 characters in length.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	// Two digit code that specifies the country in which the certificate subject located. Must be less than or equal to 2 characters in length.
	Country string `json:"country,omitempty" yaml:"country,omitempty"`

	// Disambiguating information for the certificate subject. Must be less than or equal to 64 characters in length.
	DistinguishedNameQualifier string `json:"distinguishedNameQualifier,omitempty" yaml:"distinguishedNameQualifier,omitempty"`

	// Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third. Must be less than or equal to 3 characters in length.
	GenerationQualifier string `json:"generationQualifier,omitempty" yaml:"generationQualifier,omitempty"`

	// Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`. Must be less than or equal to 5 characters in length.
	Initials string `json:"initials,omitempty" yaml:"initials,omitempty"`

	// Locality (such as a city or town) in which the certificate subject is located. Must be less than or equal to 128 characters in length.
	Locality string `json:"locality,omitempty" yaml:"locality,omitempty"`

	// Legal name of the organization with which the certificate subject is affiliated. Must be less than or equal to 64 characters in length.
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	// Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first. Must be less than or equal to 40 characters in length.
	Surname string `json:"surname,omitempty" yaml:"surname,omitempty"`

	// First name. Must be less than or equal to 16 characters in length.
	GivenName string `json:"givenName,omitempty" yaml:"givenName,omitempty"`

	// Subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated. Must be less than or equal to 64 characters in length.
	OrganizationalUnit string `json:"organizationalUnit,omitempty" yaml:"organizationalUnit,omitempty"`

	// Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza. Must be less than or equal to 128 characters in length.
	Pseudonym string `json:"pseudonym,omitempty" yaml:"pseudonym,omitempty"`

	// State in which the subject of the certificate is located. Must be less than or equal to 128 characters in length.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject. Must be less than or equal to 64 characters in length.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
