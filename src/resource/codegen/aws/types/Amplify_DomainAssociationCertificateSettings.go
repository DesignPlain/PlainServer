package types

type Amplify_DomainAssociationCertificateSettings struct {
	// DNS records for certificate verification in a space-delimited format (`<record> CNAME <target>`).
	CertificateVerificationDnsRecord string `json:"certificateVerificationDnsRecord,omitempty" yaml:"certificateVerificationDnsRecord,omitempty"`

	// The Amazon resource name (ARN) for the custom certificate.
	CustomCertificateArn string `json:"customCertificateArn,omitempty" yaml:"customCertificateArn,omitempty"`

	// The certificate type. Valid values are `AMPLIFY_MANAGED` and `CUSTOM`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
