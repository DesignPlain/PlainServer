package types

type Budgets_getBudgetNotification struct {
	// (Required) Comparison operator to use to evaluate the condition. Can be `LESS_THAN`, `EQUAL_TO` or `GREATER_THAN`.
	ComparisonOperator string `json:"comparisonOperator,omitempty" yaml:"comparisonOperator,omitempty"`

	// (Required) What kind of budget value to notify on. Can be `ACTUAL` or `FORECASTED`.
	NotificationType string `json:"notificationType,omitempty" yaml:"notificationType,omitempty"`

	// (Optional) E-Mail addresses to notify. Either this or `subscriber_sns_topic_arns` is required.
	SubscriberEmailAddresses []string `json:"subscriberEmailAddresses,omitempty" yaml:"subscriberEmailAddresses,omitempty"`

	// (Optional) SNS topics to notify. Either this or `subscriber_email_addresses` is required.
	SubscriberSnsTopicArns []string `json:"subscriberSnsTopicArns,omitempty" yaml:"subscriberSnsTopicArns,omitempty"`

	// (Required) Threshold when the notification should be sent.
	Threshold float64 `json:"threshold,omitempty" yaml:"threshold,omitempty"`

	// (Required) What kind of threshold is defined. Can be `PERCENTAGE` OR `ABSOLUTE_VALUE`.
	ThresholdType string `json:"thresholdType,omitempty" yaml:"thresholdType,omitempty"`
}
