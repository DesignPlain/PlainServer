package types

type Route53_ZoneVpc struct {
	// ID of the VPC to associate.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Region of the VPC to associate. Defaults to AWS provider region.
	VpcRegion string `json:"vpcRegion,omitempty" yaml:"vpcRegion,omitempty"`
}
