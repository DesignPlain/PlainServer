package types

type Acm_CertificateDomainValidationOption struct {
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The name of the DNS record to create to validate the certificate
	ResourceRecordName string `json:"resourceRecordName,omitempty" yaml:"resourceRecordName,omitempty"`

	// The type of DNS record to create
	ResourceRecordType string `json:"resourceRecordType,omitempty" yaml:"resourceRecordType,omitempty"`

	// The value the DNS record needs to have
	ResourceRecordValue string `json:"resourceRecordValue,omitempty" yaml:"resourceRecordValue,omitempty"`
}
