package types

type Certificatemanager_CertificateMapGclbTargetIpConfig struct {
	// An external IP address
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// A list of ports
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`
}
