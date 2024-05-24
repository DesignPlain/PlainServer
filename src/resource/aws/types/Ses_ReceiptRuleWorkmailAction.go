package types

type Ses_ReceiptRuleWorkmailAction struct {
	// The ARN of the WorkMail organization
	OrganizationArn string `json:"organizationArn,omitempty" yaml:"organizationArn,omitempty"`

	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// The ARN of an SNS topic to notify
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
