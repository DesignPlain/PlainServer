package types

type Worklink_FleetNetwork struct {
	/*
	   A list of security group IDs associated with access to the provided subnets.

	   --identity_provider-- requires the following:

	   > --NOTE:-- `identity_provider` cannot be removed without force recreating.
	*/
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of subnet IDs used for X-ENI connections from Amazon WorkLink rendering containers.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The VPC ID with connectivity to associated websites.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
