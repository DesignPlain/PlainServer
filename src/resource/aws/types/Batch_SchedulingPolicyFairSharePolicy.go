package types

type Batch_SchedulingPolicyFairSharePolicy struct {
	// A value used to reserve some of the available maximum vCPU for fair share identifiers that have not yet been used. For more information, see [FairsharePolicy](https://docs.aws.amazon.com/batch/latest/APIReference/API_FairsharePolicy.html).
	ComputeReservation int `json:"computeReservation,omitempty" yaml:"computeReservation,omitempty"`

	//
	ShareDecaySeconds int `json:"shareDecaySeconds,omitempty" yaml:"shareDecaySeconds,omitempty"`

	// One or more share distribution blocks which define the weights for the fair share identifiers for the fair share policy. For more information, see [FairsharePolicy](https://docs.aws.amazon.com/batch/latest/APIReference/API_FairsharePolicy.html). The `share_distribution` block is documented below.
	ShareDistributions []Batch_SchedulingPolicyFairSharePolicyShareDistribution `json:"shareDistributions,omitempty" yaml:"shareDistributions,omitempty"`
}
