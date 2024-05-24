package lb

type ListenerCertificate struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// The ARN of the listener to which to attach the certificate.
	ListenerArn string `json:"listenerArn,omitempty" yaml:"listenerArn,omitempty"`
}
