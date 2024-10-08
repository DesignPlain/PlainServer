package types

type Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubjectAltName struct {
	// Contains only valid RFC 3986 URIs.
	Uris []string `json:"uris,omitempty" yaml:"uris,omitempty"`

	/*
	   (Output)
	   Contains additional subject alternative name values.
	   Structure is documented below.
	*/
	CustomSans []Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSan `json:"customSans,omitempty" yaml:"customSans,omitempty"`

	// Contains only valid, fully-qualified host names.
	DnsNames []string `json:"dnsNames,omitempty" yaml:"dnsNames,omitempty"`

	// Contains only valid RFC 2822 E-mail addresses.
	EmailAddresses []string `json:"emailAddresses,omitempty" yaml:"emailAddresses,omitempty"`

	// Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`
}
