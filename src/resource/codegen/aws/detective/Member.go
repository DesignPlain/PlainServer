package detective

type Member struct {
	// Email address for the account.
	EmailAddress string `json:"emailAddress,omitempty" yaml:"emailAddress,omitempty"`

	// ARN of the behavior graph to invite the member accounts to contribute their data to.
	GraphArn string `json:"graphArn,omitempty" yaml:"graphArn,omitempty"`

	// A custom message to include in the invitation. Amazon Detective adds this message to the standard content that it sends for an invitation.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// AWS account ID for the account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// If set to true, then the root user of the invited account will _not_ receive an email notification. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. By default, this is set to `false`.
	DisableEmailNotification bool `json:"disableEmailNotification,omitempty" yaml:"disableEmailNotification,omitempty"`
}
