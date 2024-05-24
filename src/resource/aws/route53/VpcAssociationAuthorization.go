package route53

type VpcAssociationAuthorization struct {
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId string `json:"zoneId,omitempty" yaml:"zoneId,omitempty"`

	// The VPC to authorize for association with the private hosted zone.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion string `json:"vpcRegion,omitempty" yaml:"vpcRegion,omitempty"`
}
