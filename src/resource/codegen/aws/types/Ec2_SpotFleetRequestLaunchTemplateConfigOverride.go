package types

type Ec2_SpotFleetRequestLaunchTemplateConfigOverride struct {
	// The availability zone in which to place the request.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The instance requirements. See below.
	InstanceRequirements Ec2_SpotFleetRequestLaunchTemplateConfigOverrideInstanceRequirements `json:"instanceRequirements,omitempty" yaml:"instanceRequirements,omitempty"`

	// The type of instance to request.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The priority for the launch template override. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority.
	Priority float64 `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The maximum spot bid for this override request.
	SpotPrice string `json:"spotPrice,omitempty" yaml:"spotPrice,omitempty"`

	// The subnet in which to launch the requested instance.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The capacity added to the fleet by a fulfilled request.
	WeightedCapacity float64 `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`
}
