package types

type Budgets_BudgetAutoAdjustData struct {
	//
	AutoAdjustType string `json:"autoAdjustType,omitempty" yaml:"autoAdjustType,omitempty"`

	//
	HistoricalOptions Budgets_BudgetAutoAdjustDataHistoricalOptions `json:"historicalOptions,omitempty" yaml:"historicalOptions,omitempty"`

	//
	LastAutoAdjustTime string `json:"lastAutoAdjustTime,omitempty" yaml:"lastAutoAdjustTime,omitempty"`
}
