package types

type Batch_getSchedulingPolicyFairSharePolicyShareDistribution struct {
	// Fair share identifier or fair share identifier prefix. For more information, see [ShareAttributes](https://docs.aws.amazon.com/batch/latest/APIReference/API_ShareAttributes.html).
	ShareIdentifier string `json:"shareIdentifier,omitempty" yaml:"shareIdentifier,omitempty"`

	// Weight factor for the fair share identifier. For more information, see [ShareAttributes](https://docs.aws.amazon.com/batch/latest/APIReference/API_ShareAttributes.html).
	WeightFactor float64 `json:"weightFactor,omitempty" yaml:"weightFactor,omitempty"`
}
