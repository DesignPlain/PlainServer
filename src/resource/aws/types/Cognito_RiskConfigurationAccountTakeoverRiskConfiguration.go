package types

type Cognito_RiskConfigurationAccountTakeoverRiskConfiguration struct {
	// The notify configuration used to construct email notifications. See details below.
	NotifyConfiguration Cognito_RiskConfigurationAccountTakeoverRiskConfigurationNotifyConfiguration `json:"notifyConfiguration,omitempty" yaml:"notifyConfiguration,omitempty"`

	// Account takeover risk configuration actions. See details below.
	Actions Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActions `json:"actions,omitempty" yaml:"actions,omitempty"`
}
