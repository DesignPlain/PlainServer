package types

type Ec2_FleetSpotOptionsMaintenanceStrategies struct {
	// Nested argument containing the capacity rebalance for your fleet request. Defined below.
	CapacityRebalance Ec2_FleetSpotOptionsMaintenanceStrategiesCapacityRebalance `json:"capacityRebalance,omitempty" yaml:"capacityRebalance,omitempty"`
}
