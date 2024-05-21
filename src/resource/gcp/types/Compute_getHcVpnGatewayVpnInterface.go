package types

type Compute_getHcVpnGatewayVpnInterface struct {
	/*
	   URL of the interconnect attachment resource. When the value
	   of this field is present, the VPN Gateway will be used for
	   IPsec-encrypted Cloud Interconnect; all Egress or Ingress
	   traffic for this VPN Gateway interface will go through the
	   specified interconnect attachment resource.

	   Not currently available publicly.
	*/
	InterconnectAttachment string `json:"interconnectAttachment,omitempty" yaml:"interconnectAttachment,omitempty"`

	// The external IP address for this VPN gateway interface.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// The numeric ID of this VPN gateway interface.
	Id int `json:"id,omitempty" yaml:"id,omitempty"`
}
