package types

type Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfiguration struct {
	// ARN of the imported certificate authority (CA) certificate within Certificate Manager (ACM) to use for outbound SSL/TLS inspection. See [Using SSL/TLS certificates with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html) for limitations on CA certificates.
	CertificateAuthorityArn string `json:"certificateAuthorityArn,omitempty" yaml:"certificateAuthorityArn,omitempty"`

	// Check Certificate Revocation Status block. Detailed below.
	CheckCertificateRevocationStatus Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationCheckCertificateRevocationStatus `json:"checkCertificateRevocationStatus,omitempty" yaml:"checkCertificateRevocationStatus,omitempty"`

	// Scope block. Detailed below.
	Scopes []Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScope `json:"scopes,omitempty" yaml:"scopes,omitempty"`

	// Server certificates to use for inbound SSL/TLS inspection. See [Using SSL/TLS certificates with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html).
	ServerCertificates []Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationServerCertificate `json:"serverCertificates,omitempty" yaml:"serverCertificates,omitempty"`
}
