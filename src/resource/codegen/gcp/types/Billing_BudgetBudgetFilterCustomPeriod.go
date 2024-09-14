package types

type Billing_BudgetBudgetFilterCustomPeriod struct {
	/*
	   Optional. The end date of the time period. Budgets with elapsed end date won't be processed.
	   If unset, specifies to track all usage incurred since the startDate.
	   Structure is documented below.
	*/
	EndDate Billing_BudgetBudgetFilterCustomPeriodEndDate `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	/*
	   A start date is required. The start date must be after January 1, 2017.
	   Structure is documented below.
	*/
	StartDate Billing_BudgetBudgetFilterCustomPeriodStartDate `json:"startDate,omitempty" yaml:"startDate,omitempty"`
}
