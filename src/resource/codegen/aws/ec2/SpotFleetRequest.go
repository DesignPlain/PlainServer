package ec2

import types "libds/aws/types"

type SpotFleetRequest struct {
	/*
	   The number of Spot pools across which to allocate your target Spot capacity.
	   Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
	   the cheapest Spot pools and evenly allocates your target Spot capacity across
	   the number of Spot pools that you specify.
	*/
	InstancePoolsToUseCount int `json:"instancePoolsToUseCount,omitempty" yaml:"instancePoolsToUseCount,omitempty"`

	// The unit for the target capacity. This can only be done with `instance_requirements` defined
	TargetCapacityUnitType string `json:"targetCapacityUnitType,omitempty" yaml:"targetCapacityUnitType,omitempty"`

	// A list of `aws.alb.TargetGroup` ARNs, for use with Application Load Balancing.
	TargetGroupArns []string `json:"targetGroupArns,omitempty" yaml:"targetGroupArns,omitempty"`

	/*
	   Indicates whether running Spot
	   instances should be terminated when the resource is deleted (and the Spot fleet request cancelled).
	   If no value is specified, the value of the `terminate_instances_with_expiration` argument is used.
	*/
	TerminateInstancesOnDelete string `json:"terminateInstancesOnDelete,omitempty" yaml:"terminateInstancesOnDelete,omitempty"`

	/*
	   Indicates whether running Spot
	   instances should be terminated when the Spot fleet request expires.
	*/
	TerminateInstancesWithExpiration bool `json:"terminateInstancesWithExpiration,omitempty" yaml:"terminateInstancesWithExpiration,omitempty"`

	/*
	   If set, this provider will
	   wait for the Spot Request to be fulfilled, and will throw an error if the
	   timeout of 10m is reached.
	*/
	WaitForFulfillment bool `json:"waitForFulfillment,omitempty" yaml:"waitForFulfillment,omitempty"`

	// Reserved.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`

	/*
	   The type of fleet request. Indicates whether the Spot Fleet only requests the target
	   capacity or also attempts to maintain it. Default is `maintain`.
	*/
	FleetType string `json:"fleetType,omitempty" yaml:"fleetType,omitempty"`

	/*
	   Indicates whether a Spot
	   instance stops or terminates when it is interrupted. Default is
	   `terminate`.
	*/
	InstanceInterruptionBehaviour string `json:"instanceInterruptionBehaviour,omitempty" yaml:"instanceInterruptionBehaviour,omitempty"`

	/*
	   Used to define the launch configuration of the
	   spot-fleet request. Can be specified multiple times to define different bids
	   across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.

	   --Note--: This takes in similar but not
	   identical inputs as `aws.ec2.Instance`.  There are limitations on
	   what you can specify. See the list of officially supported inputs in the
	   [reference documentation](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetLaunchSpecification.html). Any normal `aws.ec2.Instance` parameter that corresponds to those inputs may be used and it have
	   a additional parameter `iam_instance_profile_arn` takes `aws.iam.InstanceProfile` attribute `arn` as input.
	*/
	LaunchSpecifications []types.Ec2_SpotFleetRequestLaunchSpecification `json:"launchSpecifications,omitempty" yaml:"launchSpecifications,omitempty"`

	// Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.
	LaunchTemplateConfigs []types.Ec2_SpotFleetRequestLaunchTemplateConfig `json:"launchTemplateConfigs,omitempty" yaml:"launchTemplateConfigs,omitempty"`

	// A list of elastic load balancer names to add to the Spot fleet.
	LoadBalancers []string `json:"loadBalancers,omitempty" yaml:"loadBalancers,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
	ValidUntil string `json:"validUntil,omitempty" yaml:"validUntil,omitempty"`

	/*
	   Indicates how to allocate the target capacity across
	   the Spot pools specified by the Spot fleet request. Valid values: `lowestPrice`, `diversified`, `capacityOptimized`, `capacityOptimizedPrioritized`, and `priceCapacityOptimized`. The default is
	   `lowestPrice`.
	*/
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`

	/*
	   Grants the Spot fleet permission to terminate
	   Spot instances on your behalf when you cancel its Spot fleet request using
	   CancelSpotFleetRequests or when the Spot fleet request expires, if you set
	   terminateInstancesWithExpiration.
	*/
	IamFleetRole string `json:"iamFleetRole,omitempty" yaml:"iamFleetRole,omitempty"`

	// The order of the launch template overrides to use in fulfilling On-Demand capacity. the possible values are: `lowestPrice` and `prioritized`. the default is `lowestPrice`.
	OnDemandAllocationStrategy string `json:"onDemandAllocationStrategy,omitempty" yaml:"onDemandAllocationStrategy,omitempty"`

	// The maximum amount per hour for On-Demand Instances that you're willing to pay. When the maximum amount you're willing to pay is reached, the fleet stops launching instances even if it hasn’t met the target capacity.
	OnDemandMaxTotalPrice string `json:"onDemandMaxTotalPrice,omitempty" yaml:"onDemandMaxTotalPrice,omitempty"`

	// Nested argument containing maintenance strategies for managing your Spot Instances that are at an elevated risk of being interrupted. Defined below.
	SpotMaintenanceStrategies types.Ec2_SpotFleetRequestSpotMaintenanceStrategies `json:"spotMaintenanceStrategies,omitempty" yaml:"spotMaintenanceStrategies,omitempty"`

	// The maximum bid price per unit hour.
	SpotPrice string `json:"spotPrice,omitempty" yaml:"spotPrice,omitempty"`

	// The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
	ValidFrom string `json:"validFrom,omitempty" yaml:"validFrom,omitempty"`

	/*
	   Indicates whether running Spot
	   instances should be terminated if the target capacity of the Spot fleet
	   request is decreased below the current size of the Spot fleet.
	*/
	ExcessCapacityTerminationPolicy string `json:"excessCapacityTerminationPolicy,omitempty" yaml:"excessCapacityTerminationPolicy,omitempty"`

	// The number of On-Demand units to request. If the request type is `maintain`, you can specify a target capacity of 0 and add capacity later.
	OnDemandTargetCapacity int `json:"onDemandTargetCapacity,omitempty" yaml:"onDemandTargetCapacity,omitempty"`

	// Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
	ReplaceUnhealthyInstances bool `json:"replaceUnhealthyInstances,omitempty" yaml:"replaceUnhealthyInstances,omitempty"`

	/*
	   The number of units to request. You can choose to set the
	   target capacity in terms of instances or a performance characteristic that is
	   important to your application workload, such as vCPUs, memory, or I/O.
	*/
	TargetCapacity int `json:"targetCapacity,omitempty" yaml:"targetCapacity,omitempty"`
}
