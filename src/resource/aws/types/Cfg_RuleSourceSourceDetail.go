package types

type Cfg_RuleSourceSourceDetail struct {
	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWSresources. This defaults to `aws.config` and is the only valid value.
	EventSource string `json:"eventSource,omitempty" yaml:"eventSource,omitempty"`

	// The frequency that you want AWS Config to run evaluations for a rule that istriggered periodically. If specified, requires `message_type` to be `ScheduledNotification`.
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" yaml:"maximumExecutionFrequency,omitempty"`

	// The type of notification that triggers AWS Config to run an evaluation for a rule. You canspecify the following notification types:
	MessageType string `json:"messageType,omitempty" yaml:"messageType,omitempty"`
}
