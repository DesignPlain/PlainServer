package billing

import types "DesignSphere_Server/src/resource/gcp/types"

type Budget struct {
	/*
	   Defines notifications that are sent on every update to the
	   billing account's spend, regardless of the thresholds defined
	   using threshold rules.
	   Structure is documented below.
	*/
	AllUpdatesRule types.Billing_BudgetAllUpdatesRule `json:"allUpdatesRule,omitempty" yaml:"allUpdatesRule,omitempty"`

	/*
	   The budgeted amount for each usage period.
	   Structure is documented below.
	*/
	Amount types.Billing_BudgetAmount `json:"amount,omitempty" yaml:"amount,omitempty"`

	// ID of the billing account to set a budget on.
	BillingAccount string `json:"billingAccount,omitempty" yaml:"billingAccount,omitempty"`

	/*
	   Filters that define which resources are used to compute the actual
	   spend against the budget.
	   Structure is documented below.
	*/
	BudgetFilter types.Billing_BudgetBudgetFilter `json:"budgetFilter,omitempty" yaml:"budgetFilter,omitempty"`

	// User data for display name in UI. Must be <= 60 chars.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Rules that trigger alerts (notifications of thresholds being
	   crossed) when spend exceeds the specified percentages of the
	   budget.
	   Structure is documented below.
	*/
	ThresholdRules []types.Billing_BudgetThresholdRule `json:"thresholdRules,omitempty" yaml:"thresholdRules,omitempty"`
}
