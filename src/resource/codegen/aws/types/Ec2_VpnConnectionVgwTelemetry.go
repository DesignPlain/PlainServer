package types

type Ec2_VpnConnectionVgwTelemetry struct {
	// The number of accepted routes.
	AcceptedRouteCount int `json:"acceptedRouteCount,omitempty" yaml:"acceptedRouteCount,omitempty"`

	// The Amazon Resource Name (ARN) of the VPN tunnel endpoint certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// The date and time of the last change in status.
	LastStatusChange string `json:"lastStatusChange,omitempty" yaml:"lastStatusChange,omitempty"`

	// The Internet-routable IP address of the virtual private gateway's outside interface.
	OutsideIpAddress string `json:"outsideIpAddress,omitempty" yaml:"outsideIpAddress,omitempty"`

	// The status of the VPN tunnel.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// If an error occurs, a description of the error.
	StatusMessage string `json:"statusMessage,omitempty" yaml:"statusMessage,omitempty"`
}
