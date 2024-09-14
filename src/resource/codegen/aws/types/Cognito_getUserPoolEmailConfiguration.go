package types

type Cognito_getUserPoolEmailConfiguration struct {
	// - Email sender address.
	From string `json:"from,omitempty" yaml:"from,omitempty"`

	// - Reply-to email address.
	ReplyToEmailAddress string `json:"replyToEmailAddress,omitempty" yaml:"replyToEmailAddress,omitempty"`

	// - Source Amazon Resource Name (ARN) for emails.
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`

	// - Configuration set used for sending emails.
	ConfigurationSet string `json:"configurationSet,omitempty" yaml:"configurationSet,omitempty"`

	// - Email sending account.
	EmailSendingAccount string `json:"emailSendingAccount,omitempty" yaml:"emailSendingAccount,omitempty"`
}
