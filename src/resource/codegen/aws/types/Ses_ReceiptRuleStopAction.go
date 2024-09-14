package types

type Ses_ReceiptRuleStopAction struct {
	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// The scope to apply. The only acceptable value is `RuleSet`.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// The ARN of an SNS topic to notify
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
