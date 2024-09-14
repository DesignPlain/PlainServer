package iam

type SigningCertificate struct {
	// The contents of the signing certificate in PEM-encoded format.
	CertificateBody string `json:"certificateBody,omitempty" yaml:"certificateBody,omitempty"`

	// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The name of the user the signing certificate is for.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
