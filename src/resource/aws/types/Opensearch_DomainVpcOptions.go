package types

type Opensearch_DomainVpcOptions struct {
	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// List of VPC Security Group IDs to be applied to the OpenSearch domain endpoints. If omitted, the default Security Group for the VPC will be used.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of VPC Subnet IDs for the OpenSearch domain endpoints to be created in.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	//
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
