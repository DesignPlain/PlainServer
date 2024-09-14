package types

type Ec2_SpotFleetRequestSpotMaintenanceStrategies struct {
	// Nested argument containing the capacity rebalance for your fleet request. Defined below.
	CapacityRebalance Ec2_SpotFleetRequestSpotMaintenanceStrategiesCapacityRebalance `json:"capacityRebalance,omitempty" yaml:"capacityRebalance,omitempty"`
}
