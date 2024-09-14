package types

type Ec2_FleetLaunchTemplateConfigOverride struct {
	// Availability Zone in which to launch the instances.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Override the instance type in the Launch Template with instance types that satisfy the requirements.
	InstanceRequirements Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirements `json:"instanceRequirements,omitempty" yaml:"instanceRequirements,omitempty"`

	// Instance type.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Maximum price per unit hour that you are willing to pay for a Spot Instance.
	MaxPrice string `json:"maxPrice,omitempty" yaml:"maxPrice,omitempty"`

	// Priority for the launch template override. If `on_demand_options` `allocation_strategy` is set to `prioritized`, EC2 Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity. The highest priority is launched first. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. Valid values are whole numbers starting at 0.
	Priority float64 `json:"priority,omitempty" yaml:"priority,omitempty"`

	// ID of the subnet in which to launch the instances.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// Number of units provided by the specified instance type.
	WeightedCapacity float64 `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`
}
