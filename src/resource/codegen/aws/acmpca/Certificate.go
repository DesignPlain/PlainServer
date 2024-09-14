package acmpca

import types "libds/aws/types"

type Certificate struct {
	// ARN of the certificate authority.
	CertificateAuthorityArn string `json:"certificateAuthorityArn,omitempty" yaml:"certificateAuthorityArn,omitempty"`

	// Certificate Signing Request in PEM format.
	CertificateSigningRequest string `json:"certificateSigningRequest,omitempty" yaml:"certificateSigningRequest,omitempty"`

	// Algorithm to use to sign certificate requests. Valid values: `SHA256WITHRSA`, `SHA256WITHECDSA`, `SHA384WITHRSA`, `SHA384WITHECDSA`, `SHA512WITHRSA`, `SHA512WITHECDSA`.
	SigningAlgorithm string `json:"signingAlgorithm,omitempty" yaml:"signingAlgorithm,omitempty"`

	/*
	   Template to use when issuing a certificate.
	   See [ACM PCA Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html) for more information.
	*/
	TemplateArn string `json:"templateArn,omitempty" yaml:"templateArn,omitempty"`

	// Configures end of the validity period for the certificate. See validity block below.
	Validity types.Acmpca_CertificateValidity `json:"validity,omitempty" yaml:"validity,omitempty"`

	// Specifies X.509 certificate information to be included in the issued certificate. To use with API Passthrough templates
	ApiPassthrough string `json:"apiPassthrough,omitempty" yaml:"apiPassthrough,omitempty"`
}
