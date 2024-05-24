package ec2

type RouteTableAssociation struct {
	// The subnet ID to create an association. Conflicts with `gateway_id`.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The gateway ID to create an association. Conflicts with `subnet_id`.
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	// The ID of the routing table to associate with.
	RouteTableId string `json:"routeTableId,omitempty" yaml:"routeTableId,omitempty"`
}
