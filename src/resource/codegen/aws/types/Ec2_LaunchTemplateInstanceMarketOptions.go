package types

type Ec2_LaunchTemplateInstanceMarketOptions struct {
	// The market type. Can be `spot`.
	MarketType string `json:"marketType,omitempty" yaml:"marketType,omitempty"`

	// The options for [Spot Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
	SpotOptions Ec2_LaunchTemplateInstanceMarketOptionsSpotOptions `json:"spotOptions,omitempty" yaml:"spotOptions,omitempty"`
}
