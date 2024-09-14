package types

type Ses_ReceiptRuleSnsAction struct {
	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// The ARN of an SNS topic to notify
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`

	// The encoding to use for the email within the Amazon SNS notification. Default value is `UTF-8`.
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}
