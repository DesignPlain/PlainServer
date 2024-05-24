package types

type Budgets_getBudgetCostFilter struct {
	/*
	   The name of a budget. Unique within accounts.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
