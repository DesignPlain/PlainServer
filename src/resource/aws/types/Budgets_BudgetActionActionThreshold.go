package types

type Budgets_BudgetActionActionThreshold struct {
	// The type of threshold for a notification. Valid values are `PERCENTAGE` or `ABSOLUTE_VALUE`.
	ActionThresholdType string `json:"actionThresholdType,omitempty" yaml:"actionThresholdType,omitempty"`

	// The threshold of a notification.
	ActionThresholdValue float64 `json:"actionThresholdValue,omitempty" yaml:"actionThresholdValue,omitempty"`
}
