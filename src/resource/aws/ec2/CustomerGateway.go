package ec2

type CustomerGateway struct {
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn string `json:"bgpAsn,omitempty" yaml:"bgpAsn,omitempty"`

	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// A name for the customer gateway device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// The IPv4 address for the customer gateway device's outside interface.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// Tags to apply to the gateway. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The type of customer gateway. The only type AWS
	   supports at this time is "ipsec.1".
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
