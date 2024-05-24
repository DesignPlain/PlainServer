package types

type Budgets_getBudgetAutoAdjustDataHistoricalOption struct {
	// (Required) - The number of budget periods included in the moving-average calculation that determines your auto-adjusted budget amount.
	BudgetAdjustmentPeriod int `json:"budgetAdjustmentPeriod,omitempty" yaml:"budgetAdjustmentPeriod,omitempty"`

	// (Optional) - The integer that describes how many budget periods in your BudgetAdjustmentPeriod are included in the calculation of your current budget limit. If the first budget period in your BudgetAdjustmentPeriod has no cost data, then that budget period isn’t included in the average that determines your budget limit. You can’t set your own LookBackAvailablePeriods. The value is automatically calculated from the `budget_adjustment_period` and your historical cost data.
	LookbackAvailablePeriods int `json:"lookbackAvailablePeriods,omitempty" yaml:"lookbackAvailablePeriods,omitempty"`
}
