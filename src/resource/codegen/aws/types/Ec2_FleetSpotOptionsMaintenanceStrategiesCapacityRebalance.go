package types

type Ec2_FleetSpotOptionsMaintenanceStrategiesCapacityRebalance struct {
	//
	TerminationDelay int `json:"terminationDelay,omitempty" yaml:"terminationDelay,omitempty"`

	// The replacement strategy to use. Only available for fleets of `type` set to `maintain`. Valid values: `launch`.
	ReplacementStrategy string `json:"replacementStrategy,omitempty" yaml:"replacementStrategy,omitempty"`
}
