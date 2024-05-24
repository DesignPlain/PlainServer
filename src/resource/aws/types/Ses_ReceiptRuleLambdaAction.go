package types

type Ses_ReceiptRuleLambdaAction struct {
	// The ARN of the Lambda function to invoke
	FunctionArn string `json:"functionArn,omitempty" yaml:"functionArn,omitempty"`

	// `Event` or `RequestResponse`
	InvocationType string `json:"invocationType,omitempty" yaml:"invocationType,omitempty"`

	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// The ARN of an SNS topic to notify
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
