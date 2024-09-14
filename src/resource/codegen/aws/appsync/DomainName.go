package appsync

type DomainName struct {
	// ARN of the certificate. This can be an Certificate Manager (ACM) certificate or an Identity and Access Management (IAM) server certificate. The certifiacte must reside in us-east-1.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// A description of the Domain Name.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Domain name.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
