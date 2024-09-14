package types

type Cognito_RiskConfigurationAccountTakeoverRiskConfigurationNotifyConfiguration struct {
	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy. This identity permits Amazon Cognito to send for the email address specified in the From parameter.
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`

	// Email template used when a detected risk event is blocked. See notify email type below.
	BlockEmail Cognito_RiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationBlockEmail `json:"blockEmail,omitempty" yaml:"blockEmail,omitempty"`

	// The email address that is sending the email. The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	From string `json:"from,omitempty" yaml:"from,omitempty"`

	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk. See notify email type below.
	MfaEmail Cognito_RiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationMfaEmail `json:"mfaEmail,omitempty" yaml:"mfaEmail,omitempty"`

	// The email template used when a detected risk event is allowed. See notify email type below.
	NoActionEmail Cognito_RiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationNoActionEmail `json:"noActionEmail,omitempty" yaml:"noActionEmail,omitempty"`

	// The destination to which the receiver of an email should reply to.
	ReplyTo string `json:"replyTo,omitempty" yaml:"replyTo,omitempty"`
}
