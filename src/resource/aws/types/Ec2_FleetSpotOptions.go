package types

type Ec2_FleetSpotOptions struct {
	// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`, `capacity-optimized`, `capacity-optimized-prioritized` and `price-capacity-optimized`. Default: `lowestPrice`.
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`

	// Behavior when a Spot Instance is interrupted. Valid values: `hibernate`, `stop`, `terminate`. Default: `terminate`.
	InstanceInterruptionBehavior string `json:"instanceInterruptionBehavior,omitempty" yaml:"instanceInterruptionBehavior,omitempty"`

	// Number of Spot pools across which to allocate your target Spot capacity. Valid only when Spot `allocation_strategy` is set to `lowestPrice`. Default: `1`.
	InstancePoolsToUseCount int `json:"instancePoolsToUseCount,omitempty" yaml:"instancePoolsToUseCount,omitempty"`

	// Nested argument containing maintenance strategies for managing your Spot Instances that are at an elevated risk of being interrupted. Defined below.
	MaintenanceStrategies Ec2_FleetSpotOptionsMaintenanceStrategies `json:"maintenanceStrategies,omitempty" yaml:"maintenanceStrategies,omitempty"`
}
