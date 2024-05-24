package route53

type ZoneAssociation struct {
	// The VPC to associate with the private hosted zone.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion string `json:"vpcRegion,omitempty" yaml:"vpcRegion,omitempty"`

	// The private hosted zone to associate.
	ZoneId string `json:"zoneId,omitempty" yaml:"zoneId,omitempty"`
}
