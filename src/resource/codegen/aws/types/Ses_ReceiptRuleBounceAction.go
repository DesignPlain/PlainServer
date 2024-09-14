package types

type Ses_ReceiptRuleBounceAction struct {
	// The RFC 3463 SMTP enhanced status code
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	// The ARN of an SNS topic to notify
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`

	// The message to send
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// The email address of the sender
	Sender string `json:"sender,omitempty" yaml:"sender,omitempty"`

	// The RFC 5321 SMTP reply code
	SmtpReplyCode string `json:"smtpReplyCode,omitempty" yaml:"smtpReplyCode,omitempty"`
}
