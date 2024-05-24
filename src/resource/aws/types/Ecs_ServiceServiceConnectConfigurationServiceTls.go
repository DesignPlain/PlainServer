package types

type Ecs_ServiceServiceConnectConfigurationServiceTls struct {
	// The details of the certificate authority which will issue the certificate.
	IssuerCertAuthority Ecs_ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthority `json:"issuerCertAuthority,omitempty" yaml:"issuerCertAuthority,omitempty"`

	// The KMS key used to encrypt the private key in Secrets Manager.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	// The ARN of the IAM Role that's associated with the Service Connect TLS.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
