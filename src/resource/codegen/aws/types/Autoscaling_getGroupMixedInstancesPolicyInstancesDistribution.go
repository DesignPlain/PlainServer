package types

type Autoscaling_getGroupMixedInstancesPolicyInstancesDistribution struct {
	//
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity,omitempty" yaml:"onDemandPercentageAboveBaseCapacity,omitempty"`

	// Strategy used when launching Spot instances.
	SpotAllocationStrategy string `json:"spotAllocationStrategy,omitempty" yaml:"spotAllocationStrategy,omitempty"`

	// Number of Spot pools per availability zone to allocate capacity.
	SpotInstancePools int `json:"spotInstancePools,omitempty" yaml:"spotInstancePools,omitempty"`

	// Maximum price per unit hour that the user is willing to pay for the Spot instances.
	SpotMaxPrice string `json:"spotMaxPrice,omitempty" yaml:"spotMaxPrice,omitempty"`

	// Strategy used when launching on-demand instances.
	OnDemandAllocationStrategy string `json:"onDemandAllocationStrategy,omitempty" yaml:"onDemandAllocationStrategy,omitempty"`

	// Absolute minimum amount of desired capacity that must be fulfilled by on-demand instances.
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity,omitempty" yaml:"onDemandBaseCapacity,omitempty"`
}
