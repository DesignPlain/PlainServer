package types

type Budgets_getBudgetBudgetLimit struct {
	// The cost or usage amount that's associated with a budget forecast, actual spend, or budget threshold. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
	Amount string `json:"amount,omitempty" yaml:"amount,omitempty"`

	// The unit of measurement that's used for the budget forecast, actual spend, or budget threshold, such as USD or GBP. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
