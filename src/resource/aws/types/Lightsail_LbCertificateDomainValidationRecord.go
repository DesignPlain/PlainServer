package types

type Lightsail_LbCertificateDomainValidationRecord struct {
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	//
	ResourceRecordName string `json:"resourceRecordName,omitempty" yaml:"resourceRecordName,omitempty"`

	//
	ResourceRecordType string `json:"resourceRecordType,omitempty" yaml:"resourceRecordType,omitempty"`

	//
	ResourceRecordValue string `json:"resourceRecordValue,omitempty" yaml:"resourceRecordValue,omitempty"`
}
