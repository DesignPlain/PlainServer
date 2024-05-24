package types

type Ec2_FleetOnDemandOptions struct {
	/*
	   The minimum target capacity for On-Demand Instances in the fleet. If the minimum target capacity is not reached, the fleet launches no instances. Supported only for fleets of type `instant`.
	   If you specify `min_target_capacity`, at least one of the following must be specified: `single_availability_zone` or `single_instance_type`.
	*/
	MinTargetCapacity int `json:"minTargetCapacity,omitempty" yaml:"minTargetCapacity,omitempty"`

	// Indicates that the fleet launches all On-Demand Instances into a single Availability Zone. Supported only for fleets of type `instant`.
	SingleAvailabilityZone bool `json:"singleAvailabilityZone,omitempty" yaml:"singleAvailabilityZone,omitempty"`

	// Indicates that the fleet uses a single instance type to launch all On-Demand Instances in the fleet. Supported only for fleets of type `instant`.
	SingleInstanceType bool `json:"singleInstanceType,omitempty" yaml:"singleInstanceType,omitempty"`

	// The order of the launch template overrides to use in fulfilling On-Demand capacity. Valid values: `lowestPrice`, `prioritized`. Default: `lowestPrice`.
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`

	// The maximum amount per hour for On-Demand Instances that you're willing to pay.
	MaxTotalPrice string `json:"maxTotalPrice,omitempty" yaml:"maxTotalPrice,omitempty"`
}
