package iot

import types "DesignSphere_Server/src/resource/aws/types"

type CaCertificate struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   PEM encoded verification certificate containing the common name of a registration code. Review
	   [CreateVerificationCSR](https://docs.aws.amazon.com/iot/latest/developerguide/register-CA-cert.html). Reuired if `certificate_mode` is `DEFAULT`.
	*/
	VerificationCertificatePem string `json:"verificationCertificatePem,omitempty" yaml:"verificationCertificatePem,omitempty"`

	// Boolean flag to indicate if the certificate should be active for device authentication.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Boolean flag to indicate if the certificate should be active for device regisration.
	AllowAutoRegistration bool `json:"allowAutoRegistration,omitempty" yaml:"allowAutoRegistration,omitempty"`

	// PEM encoded CA certificate.
	CaCertificatePem string `json:"caCertificatePem,omitempty" yaml:"caCertificatePem,omitempty"`

	// The certificate mode in which the CA will be registered. Valida values: `DEFAULT` and `SNI_ONLY`. Default: `DEFAULT`.
	CertificateMode string `json:"certificateMode,omitempty" yaml:"certificateMode,omitempty"`

	// Information about the registration configuration. See below.
	RegistrationConfig types.Iot_CaCertificateRegistrationConfig `json:"registrationConfig,omitempty" yaml:"registrationConfig,omitempty"`
}
