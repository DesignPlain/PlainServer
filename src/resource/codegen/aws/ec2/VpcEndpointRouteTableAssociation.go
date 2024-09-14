package ec2

type VpcEndpointRouteTableAssociation struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId string `json:"routeTableId,omitempty" yaml:"routeTableId,omitempty"`

	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`
}
