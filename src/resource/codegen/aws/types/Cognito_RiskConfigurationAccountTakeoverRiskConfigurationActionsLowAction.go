package types

type Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction struct {
	//
	EventAction string `json:"eventAction,omitempty" yaml:"eventAction,omitempty"`

	// Whether to send a notification.
	Notify bool `json:"notify,omitempty" yaml:"notify,omitempty"`
}
