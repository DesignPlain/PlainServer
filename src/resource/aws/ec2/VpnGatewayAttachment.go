package ec2

type VpnGatewayAttachment struct {
	// The ID of the Virtual Private Gateway.
	VpnGatewayId string `json:"vpnGatewayId,omitempty" yaml:"vpnGatewayId,omitempty"`

	// The ID of the VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
