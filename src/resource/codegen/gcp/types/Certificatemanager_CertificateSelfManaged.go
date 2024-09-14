package types

type Certificatemanager_CertificateSelfManaged struct {
	/*
	   (Optional, Deprecated)
	   The certificate chain in PEM-encoded form.
	   Leaf certificate comes first, followed by intermediate ones if any.
	   --Note--: This property is sensitive and will not be displayed in the plan.

	   > --Warning:-- `certificate_pem` is deprecated and will be removed in a future major release. Use `pem_certificate` instead.
	*/
	CertificatePem string `json:"certificatePem,omitempty" yaml:"certificatePem,omitempty"`

	/*
	   The certificate chain in PEM-encoded form.
	   Leaf certificate comes first, followed by intermediate ones if any.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	PemCertificate string `json:"pemCertificate,omitempty" yaml:"pemCertificate,omitempty"`

	/*
	   The private key of the leaf certificate in PEM-encoded form.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	PemPrivateKey string `json:"pemPrivateKey,omitempty" yaml:"pemPrivateKey,omitempty"`

	/*
	   (Optional, Deprecated)
	   The private key of the leaf certificate in PEM-encoded form.
	   --Note--: This property is sensitive and will not be displayed in the plan.

	   > --Warning:-- `private_key_pem` is deprecated and will be removed in a future major release. Use `pem_private_key` instead.
	*/
	PrivateKeyPem string `json:"privateKeyPem,omitempty" yaml:"privateKeyPem,omitempty"`
}
