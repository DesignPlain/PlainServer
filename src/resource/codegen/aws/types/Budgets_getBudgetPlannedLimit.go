package types

type Budgets_getBudgetPlannedLimit struct {
	// The cost or usage amount that's associated with a budget forecast, actual spend, or budget threshold. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
	Amount string `json:"amount,omitempty" yaml:"amount,omitempty"`

	// (Required) The start time of the budget limit. Format: `2017-01-01_12:00`. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// The unit of measurement that's used for the budget forecast, actual spend, or budget threshold, such as USD or GBP. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
