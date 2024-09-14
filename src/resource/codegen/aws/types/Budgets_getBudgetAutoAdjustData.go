package types

type Budgets_getBudgetAutoAdjustData struct {
	// (Required) - The string that defines whether your budget auto-adjusts based on historical or forecasted data. Valid values: `FORECAST`,`HISTORICAL`.
	AutoAdjustType string `json:"autoAdjustType,omitempty" yaml:"autoAdjustType,omitempty"`

	// (Optional) - Configuration block of Historical Options. Required for `auto_adjust_type` of `HISTORICAL` Configuration block that defines the historical data that your auto-adjusting budget is based on.
	HistoricalOptions []Budgets_getBudgetAutoAdjustDataHistoricalOption `json:"historicalOptions,omitempty" yaml:"historicalOptions,omitempty"`

	// (Optional) - The last time that your budget was auto-adjusted.
	LastAutoAdjustTime string `json:"lastAutoAdjustTime,omitempty" yaml:"lastAutoAdjustTime,omitempty"`
}
