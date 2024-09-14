package types

type Ec2_SpotFleetRequestSpotMaintenanceStrategiesCapacityRebalance struct {
	// The replacement strategy to use. Only available for spot fleets with `fleet_type` set to `maintain`. Valid values: `launch`.
	ReplacementStrategy string `json:"replacementStrategy,omitempty" yaml:"replacementStrategy,omitempty"`
}
