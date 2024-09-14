package types

type Budgets_BudgetCostFilter struct {
	// The name of a budget. Unique within accounts.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
