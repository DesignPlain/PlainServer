package types

type Ec2_getLaunchTemplateInstanceMarketOption struct {
	//
	MarketType string `json:"marketType,omitempty" yaml:"marketType,omitempty"`

	//
	SpotOptions []Ec2_getLaunchTemplateInstanceMarketOptionSpotOption `json:"spotOptions,omitempty" yaml:"spotOptions,omitempty"`
}
