package types

type Cfg_RuleSourceSourceDetail struct {
	// The frequency that you want AWS Config to run evaluations for a rule that istriggered periodically. If specified, requires `message_type` to be `ScheduledNotification`.
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" yaml:"maximumExecutionFrequency,omitempty"`

	/*
	   The type of notification that triggers AWS Config to run an evaluation for a rule. You canspecify the following notification types:
	   - `ConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers a configuration item as a result of a resource change.
	   - `OversizedConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	   - `ScheduledNotification` - Triggers a periodic evaluation at the frequency specified for `maximum_execution_frequency`.
	   - `ConfigurationSnapshotDeliveryCompleted` - Triggers a periodic evaluation when AWS Config delivers a configuration snapshot.
	*/
	MessageType string `json:"messageType,omitempty" yaml:"messageType,omitempty"`

	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWSresources. This defaults to `aws.config` and is the only valid value.
	EventSource string `json:"eventSource,omitempty" yaml:"eventSource,omitempty"`
}
