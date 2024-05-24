package ec2

type VpnGatewayRoutePropagation struct {
	// The id of the `aws.ec2.RouteTable` to propagate routes into.
	RouteTableId string `json:"routeTableId,omitempty" yaml:"routeTableId,omitempty"`

	// The id of the `aws.ec2.VpnGateway` to propagate routes from.
	VpnGatewayId string `json:"vpnGatewayId,omitempty" yaml:"vpnGatewayId,omitempty"`
}
