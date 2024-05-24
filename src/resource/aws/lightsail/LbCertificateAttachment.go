package lightsail

type LbCertificateAttachment struct {
	// The name of your SSL/TLS certificate.
	CertificateName string `json:"certificateName,omitempty" yaml:"certificateName,omitempty"`

	// The name of the load balancer to which you want to associate the SSL/TLS certificate.
	LbName string `json:"lbName,omitempty" yaml:"lbName,omitempty"`
}
