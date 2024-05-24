package types

type Lightsail_CertificateDomainValidationOption struct {
	// A domain name for which the certificate should be issued.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	//
	ResourceRecordName string `json:"resourceRecordName,omitempty" yaml:"resourceRecordName,omitempty"`

	//
	ResourceRecordType string `json:"resourceRecordType,omitempty" yaml:"resourceRecordType,omitempty"`

	//
	ResourceRecordValue string `json:"resourceRecordValue,omitempty" yaml:"resourceRecordValue,omitempty"`
}
