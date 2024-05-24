package types

type Ec2_FleetTargetCapacitySpecification struct {
	// Default target capacity type. Valid values: `on-demand`, `spot`.
	DefaultTargetCapacityType string `json:"defaultTargetCapacityType,omitempty" yaml:"defaultTargetCapacityType,omitempty"`

	// The number of On-Demand units to request.
	OnDemandTargetCapacity int `json:"onDemandTargetCapacity,omitempty" yaml:"onDemandTargetCapacity,omitempty"`

	// The number of Spot units to request.
	SpotTargetCapacity int `json:"spotTargetCapacity,omitempty" yaml:"spotTargetCapacity,omitempty"`

	/*
	   The unit for the target capacity.
	   If you specify `target_capacity_unit_type`, `instance_requirements` must be specified.
	*/
	TargetCapacityUnitType string `json:"targetCapacityUnitType,omitempty" yaml:"targetCapacityUnitType,omitempty"`

	// The number of units to request, filled using `default_target_capacity_type`.
	TotalTargetCapacity int `json:"totalTargetCapacity,omitempty" yaml:"totalTargetCapacity,omitempty"`
}
