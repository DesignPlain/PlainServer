package types

type Acmpca_CertificateAuthorityCertificateAuthorityConfiguration struct {
	// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`

	// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
	SigningAlgorithm string `json:"signingAlgorithm,omitempty" yaml:"signingAlgorithm,omitempty"`

	// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
	Subject Acmpca_CertificateAuthorityCertificateAuthorityConfigurationSubject `json:"subject,omitempty" yaml:"subject,omitempty"`
}
