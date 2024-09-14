package types

type Cognito_UserPoolAdminCreateUserConfigInviteMessageTemplate struct {
	// Message template for email messages. Must contain `{username}` and `{####}` placeholders, for username and temporary password, respectively.
	EmailMessage string `json:"emailMessage,omitempty" yaml:"emailMessage,omitempty"`

	// Subject line for email messages.
	EmailSubject string `json:"emailSubject,omitempty" yaml:"emailSubject,omitempty"`

	// Message template for SMS messages. Must contain `{username}` and `{####}` placeholders, for username and temporary password, respectively.
	SmsMessage string `json:"smsMessage,omitempty" yaml:"smsMessage,omitempty"`
}
