package types

type Autoscaling_GroupMixedInstancesPolicyInstancesDistribution struct {
	// Strategy to use when launching on-demand instances. Valid values: `prioritized`, `lowest-price`. Default: `prioritized`.
	OnDemandAllocationStrategy string `json:"onDemandAllocationStrategy,omitempty" yaml:"onDemandAllocationStrategy,omitempty"`

	// Absolute minimum amount of desired capacity that must be fulfilled by on-demand instances. Default: `0`.
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity,omitempty" yaml:"onDemandBaseCapacity,omitempty"`

	// Percentage split between on-demand and Spot instances above the base on-demand capacity. Default: `100`.
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity,omitempty" yaml:"onDemandPercentageAboveBaseCapacity,omitempty"`

	// How to allocate capacity across the Spot pools. Valid values: `lowest-price`, `capacity-optimized`, `capacity-optimized-prioritized`, and `price-capacity-optimized`. Default: `lowest-price`.
	SpotAllocationStrategy string `json:"spotAllocationStrategy,omitempty" yaml:"spotAllocationStrategy,omitempty"`

	// Number of Spot pools per availability zone to allocate capacity. EC2 Auto Scaling selects the cheapest Spot pools and evenly allocates Spot capacity across the number of Spot pools that you specify. Only available with `spot_allocation_strategy` set to `lowest-price`. Otherwise it must be set to `0`, if it has been defined before. Default: `2`.
	SpotInstancePools int `json:"spotInstancePools,omitempty" yaml:"spotInstancePools,omitempty"`

	// Maximum price per unit hour that the user is willing to pay for the Spot instances. Default: an empty string which means the on-demand price.
	SpotMaxPrice string `json:"spotMaxPrice,omitempty" yaml:"spotMaxPrice,omitempty"`
}
