package types

type Finspace_KxClusterVpcConfiguration struct {
	//
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Identifier of the VPC endpoint
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// IP address type for cluster network configuration parameters. The following type is available: IP_V4 - IP address version 4.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	/*
	   Unique identifier of the VPC security group applied to the VPC endpoint ENI for the cluster.
	   - `subnet_ids `- (Required) Identifier of the subnet that the Privatelink VPC endpoint uses to connect to the cluster.
	*/
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`
}
