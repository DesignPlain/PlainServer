package types

type Ec2_InstanceInstanceMarketOptions struct {
	// Type of market for the instance. Valid values are `spot` and `capacity-block`. Defaults to `spot`. Required if `spot_options` is specified.
	MarketType string `json:"marketType,omitempty" yaml:"marketType,omitempty"`

	// Block to configure the options for Spot Instances. See Spot Options below for details on attributes.
	SpotOptions Ec2_InstanceInstanceMarketOptionsSpotOptions `json:"spotOptions,omitempty" yaml:"spotOptions,omitempty"`
}
