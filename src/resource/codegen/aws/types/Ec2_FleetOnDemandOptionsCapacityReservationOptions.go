package types

type Ec2_FleetOnDemandOptionsCapacityReservationOptions struct {
	// Indicates whether to use unused Capacity Reservations for fulfilling On-Demand capacity. Valid values: `use-capacity-reservations-first`.
	UsageStrategy string `json:"usageStrategy,omitempty" yaml:"usageStrategy,omitempty"`
}
