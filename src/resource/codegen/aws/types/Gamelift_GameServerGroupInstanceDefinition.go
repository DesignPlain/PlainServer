package types

type Gamelift_GameServerGroupInstanceDefinition struct {
	// An EC2 instance type.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	/*
	   Instance weighting that indicates how much this instance type contributes
	   to the total capacity of a game server group.
	   Instance weights are used by GameLift FleetIQ to calculate the instance type's cost per unit hour and better identify
	   the most cost-effective options.
	*/
	WeightedCapacity string `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`
}
