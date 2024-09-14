package types

type Ec2_getRouteTableAssociation struct {
	// ID of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	// Whether the association is due to the main route table.
	Main bool `json:"main,omitempty" yaml:"main,omitempty"`

	// Association ID.
	RouteTableAssociationId string `json:"routeTableAssociationId,omitempty" yaml:"routeTableAssociationId,omitempty"`

	// ID of the specific Route Table to retrieve.
	RouteTableId string `json:"routeTableId,omitempty" yaml:"routeTableId,omitempty"`

	// ID of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
