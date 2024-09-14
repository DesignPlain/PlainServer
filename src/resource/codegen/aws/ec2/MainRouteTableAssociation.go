package ec2

type MainRouteTableAssociation struct {
	/*
	   The ID of the Route Table to set as the new
	   main route table for the target VPC
	*/
	RouteTableId string `json:"routeTableId,omitempty" yaml:"routeTableId,omitempty"`

	// The ID of the VPC whose main route table should be set
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
