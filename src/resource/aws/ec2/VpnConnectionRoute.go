package ec2

type VpnConnectionRoute struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// The ID of the VPN connection.
	VpnConnectionId string `json:"vpnConnectionId,omitempty" yaml:"vpnConnectionId,omitempty"`
}
