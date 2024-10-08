package types

type Billing_BudgetThresholdRule struct {
	/*
	   Send an alert when this threshold is exceeded. This is a
	   1.0-based percentage, so 0.5 = 50%!!(MISSING) (MISSING)Must be >= 0.
	*/
	ThresholdPercent float64 `json:"thresholdPercent,omitempty" yaml:"thresholdPercent,omitempty"`

	/*
	   The type of basis used to determine if spend has passed
	   the threshold.
	   Default value is `CURRENT_SPEND`.
	   Possible values are: `CURRENT_SPEND`, `FORECASTED_SPEND`.
	*/
	SpendBasis string `json:"spendBasis,omitempty" yaml:"spendBasis,omitempty"`
}
