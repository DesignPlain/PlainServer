package types

type Apprunner_CustomDomainAssociationCertificateValidationRecord struct {
	// Current state of the certificate CNAME record validation. It should change to `SUCCESS` after App Runner completes validation with your DNS.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Record type, always `CNAME`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Certificate CNAME record value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Certificate CNAME record name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
