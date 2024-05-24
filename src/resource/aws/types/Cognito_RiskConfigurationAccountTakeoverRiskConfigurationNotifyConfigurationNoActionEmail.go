package types

type Cognito_RiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationNoActionEmail struct {
	// The email HTML body.
	HtmlBody string `json:"htmlBody,omitempty" yaml:"htmlBody,omitempty"`

	// The email subject.
	Subject string `json:"subject,omitempty" yaml:"subject,omitempty"`

	// The email text body.
	TextBody string `json:"textBody,omitempty" yaml:"textBody,omitempty"`
}
