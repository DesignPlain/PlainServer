package types

type Opensearch_getDomainVpcOption struct {
	// Availability zones used by the domain.
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// Security groups used by the domain.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Subnets used by the domain.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// VPC used by the domain.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
