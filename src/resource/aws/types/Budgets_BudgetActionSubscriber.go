package types

type Budgets_BudgetActionSubscriber struct {
	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The type of notification that AWS sends to a subscriber. Valid values are `SNS` or `EMAIL`.
	SubscriptionType string `json:"subscriptionType,omitempty" yaml:"subscriptionType,omitempty"`
}
