package redshiftserverless

type CustomDomainAssociation struct {
	// ARN of the certificate for the custom domain association.
	CustomDomainCertificateArn string `json:"customDomainCertificateArn,omitempty" yaml:"customDomainCertificateArn,omitempty"`

	// Custom domain to associate with the workgroup.
	CustomDomainName string `json:"customDomainName,omitempty" yaml:"customDomainName,omitempty"`

	// Name of the workgroup.
	WorkgroupName string `json:"workgroupName,omitempty" yaml:"workgroupName,omitempty"`
}
