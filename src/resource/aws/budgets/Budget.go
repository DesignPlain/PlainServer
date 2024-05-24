package budgets

import types "DesignSphere_Server/src/resource/aws/types"

type Budget struct {
	// The name of a budget. Unique within accounts.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
	Notifications []types.Budgets_BudgetNotification `json:"notifications,omitempty" yaml:"notifications,omitempty"`

	// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
	PlannedLimits []types.Budgets_BudgetPlannedLimit `json:"plannedLimits,omitempty" yaml:"plannedLimits,omitempty"`

	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
	TimeUnit string `json:"timeUnit,omitempty" yaml:"timeUnit,omitempty"`

	// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
	AutoAdjustData types.Budgets_BudgetAutoAdjustData `json:"autoAdjustData,omitempty" yaml:"autoAdjustData,omitempty"`

	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd string `json:"timePeriodEnd,omitempty" yaml:"timePeriodEnd,omitempty"`

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes types.Budgets_BudgetCostTypes `json:"costTypes,omitempty" yaml:"costTypes,omitempty"`

	// The amount of cost or usage being measured for a budget.
	LimitAmount string `json:"limitAmount,omitempty" yaml:"limitAmount,omitempty"`

	// Whether this budget tracks monetary cost or usage.
	BudgetType string `json:"budgetType,omitempty" yaml:"budgetType,omitempty"`

	// A list of CostFilter name/values pair to apply to budget.
	CostFilters []types.Budgets_BudgetCostFilter `json:"costFilters,omitempty" yaml:"costFilters,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit string `json:"limitUnit,omitempty" yaml:"limitUnit,omitempty"`

	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart string `json:"timePeriodStart,omitempty" yaml:"timePeriodStart,omitempty"`
}
