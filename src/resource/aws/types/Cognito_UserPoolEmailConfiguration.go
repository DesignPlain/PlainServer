package types

type Cognito_UserPoolEmailConfiguration struct {
	// Email configuration set name from SES.
	ConfigurationSet string `json:"configurationSet,omitempty" yaml:"configurationSet,omitempty"`

	// Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration. Required to be `DEVELOPER` if `from_email_address` is set.
	EmailSendingAccount string `json:"emailSendingAccount,omitempty" yaml:"emailSendingAccount,omitempty"`

	// Sender’s email address or sender’s display name with their email address (e.g., `john@example.com`, `John Smith <john@example.com>` or `\"John Smith Ph.D.\" <john@example.com>`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
	FromEmailAddress string `json:"fromEmailAddress,omitempty" yaml:"fromEmailAddress,omitempty"`

	// REPLY-TO email address.
	ReplyToEmailAddress string `json:"replyToEmailAddress,omitempty" yaml:"replyToEmailAddress,omitempty"`

	// ARN of the SES verified email identity to use. Required if `email_sending_account` is set to `DEVELOPER`.
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`
}
