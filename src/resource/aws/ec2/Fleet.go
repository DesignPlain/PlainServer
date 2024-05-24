package ec2

import types "DesignSphere_Server/src/resource/aws/types"

type Fleet struct {
	// Nested argument containing EC2 Launch Template configurations. Defined below.
	LaunchTemplateConfigs []types.Ec2_FleetLaunchTemplateConfig `json:"launchTemplateConfigs,omitempty" yaml:"launchTemplateConfigs,omitempty"`

	// Nested argument containing On-Demand configurations. Defined below.
	OnDemandOptions types.Ec2_FleetOnDemandOptions `json:"onDemandOptions,omitempty" yaml:"onDemandOptions,omitempty"`

	// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
	ReplaceUnhealthyInstances bool `json:"replaceUnhealthyInstances,omitempty" yaml:"replaceUnhealthyInstances,omitempty"`

	// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
	TerminateInstances bool `json:"terminateInstances,omitempty" yaml:"terminateInstances,omitempty"`

	// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
	ExcessCapacityTerminationPolicy string `json:"excessCapacityTerminationPolicy,omitempty" yaml:"excessCapacityTerminationPolicy,omitempty"`

	// Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
	FleetInstanceSets []types.Ec2_FleetFleetInstanceSet `json:"fleetInstanceSets,omitempty" yaml:"fleetInstanceSets,omitempty"`

	// Nested argument containing target capacity configurations. Defined below.
	TargetCapacitySpecification types.Ec2_FleetTargetCapacitySpecification `json:"targetCapacitySpecification,omitempty" yaml:"targetCapacitySpecification,omitempty"`

	// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
	TerminateInstancesWithExpiration bool `json:"terminateInstancesWithExpiration,omitempty" yaml:"terminateInstancesWithExpiration,omitempty"`

	// The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
	ValidUntil string `json:"validUntil,omitempty" yaml:"validUntil,omitempty"`

	// Reserved.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`

	// The number of units fulfilled by this request compared to the set target On-Demand capacity.
	FulfilledOnDemandCapacity float64 `json:"fulfilledOnDemandCapacity,omitempty" yaml:"fulfilledOnDemandCapacity,omitempty"`

	// The state of the EC2 Fleet.
	FleetState string `json:"fleetState,omitempty" yaml:"fleetState,omitempty"`

	// Nested argument containing Spot configurations. Defined below.
	SpotOptions types.Ec2_FleetSpotOptions `json:"spotOptions,omitempty" yaml:"spotOptions,omitempty"`

	// The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
	ValidFrom string `json:"validFrom,omitempty" yaml:"validFrom,omitempty"`

	// The number of units fulfilled by this request compared to the set target capacity.
	FulfilledCapacity float64 `json:"fulfilledCapacity,omitempty" yaml:"fulfilledCapacity,omitempty"`

	// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
