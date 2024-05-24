package ec2

type VpcEndpointSubnetAssociation struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`
}
