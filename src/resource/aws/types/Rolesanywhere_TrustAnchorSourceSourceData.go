package types

type Rolesanywhere_TrustAnchorSourceSourceData struct {
	//
	X509CertificateData string `json:"x509CertificateData,omitempty" yaml:"x509CertificateData,omitempty"`

	// The ARN of an ACM Private Certificate Authority.
	AcmPcaArn string `json:"acmPcaArn,omitempty" yaml:"acmPcaArn,omitempty"`
}
