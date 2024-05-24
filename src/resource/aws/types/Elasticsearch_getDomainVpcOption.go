package types

type Elasticsearch_getDomainVpcOption struct {
	// The security groups used by the domain.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The subnets used by the domain.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The VPC used by the domain.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The availability zones used by the domain.
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`
}
