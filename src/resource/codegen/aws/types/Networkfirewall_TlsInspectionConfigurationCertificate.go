package types

type Networkfirewall_TlsInspectionConfigurationCertificate struct {
	// ARN of the certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Serial number of the certificate.
	CertificateSerial string `json:"certificateSerial,omitempty" yaml:"certificateSerial,omitempty"`

	// Status of the certificate.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Details about the certificate status, including information about certificate errors.
	StatusMessage string `json:"statusMessage,omitempty" yaml:"statusMessage,omitempty"`
}
