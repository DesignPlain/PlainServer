package iot

type Certificate struct {
	// The CA certificate for the certificate to be registered. If this is set, the CA needs to be registered with AWS IoT beforehand.
	CaPem string `json:"caPem,omitempty" yaml:"caPem,omitempty"`

	/*
	   The certificate to be registered. If `ca_pem` is unspecified, review
	   [RegisterCertificateWithoutCA](https://docs.aws.amazon.com/iot/latest/apireference/API_RegisterCertificateWithoutCA.html).
	   If `ca_pem` is specified, review
	   [RegisterCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_RegisterCertificate.html)
	   for more information on registering a certificate.
	*/
	CertificatePem string `json:"certificatePem,omitempty" yaml:"certificatePem,omitempty"`

	/*
	   The certificate signing request. Review
	   [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	   for more information on generating a certificate from a certificate signing request (CSR).
	   If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
	   for more information on generating keys and a certificate.
	*/
	Csr string `json:"csr,omitempty" yaml:"csr,omitempty"`

	// Boolean flag to indicate if the certificate should be active
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
}
