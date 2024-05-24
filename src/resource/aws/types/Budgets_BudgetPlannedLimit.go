package types

type Budgets_BudgetPlannedLimit struct {
	// (Required) The amount of cost or usage being measured for a budget.
	Amount string `json:"amount,omitempty" yaml:"amount,omitempty"`

	// (Required) The start time of the budget limit. Format: `2017-01-01_12:00`. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// (Required) The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
