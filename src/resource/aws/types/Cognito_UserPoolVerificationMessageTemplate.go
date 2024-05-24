package types

type Cognito_UserPoolVerificationMessageTemplate struct {
	// Default email option. Must be either `CONFIRM_WITH_CODE` or `CONFIRM_WITH_LINK`. Defaults to `CONFIRM_WITH_CODE`.
	DefaultEmailOption string `json:"defaultEmailOption,omitempty" yaml:"defaultEmailOption,omitempty"`

	// Email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
	EmailMessage string `json:"emailMessage,omitempty" yaml:"emailMessage,omitempty"`

	// Email message template for sending a confirmation link to the user, it must contain the `{##Click Here##}` placeholder.
	EmailMessageByLink string `json:"emailMessageByLink,omitempty" yaml:"emailMessageByLink,omitempty"`

	// Subject line for the email message template. Conflicts with `email_verification_subject` argument.
	EmailSubject string `json:"emailSubject,omitempty" yaml:"emailSubject,omitempty"`

	// Subject line for the email message template for sending a confirmation link to the user.
	EmailSubjectByLink string `json:"emailSubjectByLink,omitempty" yaml:"emailSubjectByLink,omitempty"`

	// SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.
	SmsMessage string `json:"smsMessage,omitempty" yaml:"smsMessage,omitempty"`
}
