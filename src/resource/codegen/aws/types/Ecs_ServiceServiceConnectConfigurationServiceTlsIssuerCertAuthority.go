package types

type Ecs_ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthority struct {
	// ARN of the `aws.acmpca.CertificateAuthority` used to create the TLS Certificates.
	AwsPcaAuthorityArn string `json:"awsPcaAuthorityArn,omitempty" yaml:"awsPcaAuthorityArn,omitempty"`
}
