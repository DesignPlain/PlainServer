package types

type Ec2_InstanceInstanceMarketOptions struct {
	// Type of market for the instance. Valid value is `spot`. Defaults to `spot`.
	MarketType string `json:"marketType,omitempty" yaml:"marketType,omitempty"`

	// Block to configure the options for Spot Instances. See Spot Options below for details on attributes.
	SpotOptions Ec2_InstanceInstanceMarketOptionsSpotOptions `json:"spotOptions,omitempty" yaml:"spotOptions,omitempty"`
}
