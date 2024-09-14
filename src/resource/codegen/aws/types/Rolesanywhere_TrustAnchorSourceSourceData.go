package types

type Rolesanywhere_TrustAnchorSourceSourceData struct {
	// The ARN of an ACM Private Certificate Authority.
	AcmPcaArn string `json:"acmPcaArn,omitempty" yaml:"acmPcaArn,omitempty"`

	//
	X509CertificateData string `json:"x509CertificateData,omitempty" yaml:"x509CertificateData,omitempty"`
}
