package types

type Emr_ClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Specifies the strategy to use in launching On-Demand instance fleets. Currently, the only option is `lowest-price` (the default), which launches the lowest price first.
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`
}
