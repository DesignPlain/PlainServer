package types

type Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction struct {
	// The action to take in response to the account takeover action. Valid values are `BLOCK`, `MFA_IF_CONFIGURED`, `MFA_REQUIRED` and `NO_ACTION`.
	EventAction string `json:"eventAction,omitempty" yaml:"eventAction,omitempty"`

	// Whether to send a notification.
	Notify bool `json:"notify,omitempty" yaml:"notify,omitempty"`
}
