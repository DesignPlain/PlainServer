package types

type Cognito_getUserPoolAdminCreateUserConfigInviteMessageTemplate struct {
	// - Email message subject.
	EmailSubject string `json:"emailSubject,omitempty" yaml:"emailSubject,omitempty"`

	// - SMS message content.
	SmsMessage string `json:"smsMessage,omitempty" yaml:"smsMessage,omitempty"`

	// - Email message content.
	EmailMessage string `json:"emailMessage,omitempty" yaml:"emailMessage,omitempty"`
}
