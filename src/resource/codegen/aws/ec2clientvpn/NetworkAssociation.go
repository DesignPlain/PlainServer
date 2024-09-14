package ec2clientvpn

type NetworkAssociation struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId string `json:"clientVpnEndpointId,omitempty" yaml:"clientVpnEndpointId,omitempty"`

	// The ID of the subnet to associate with the Client VPN endpoint.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
