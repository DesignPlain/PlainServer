package types

type Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActions struct {
	// Action to take for a high risk. See action block below.
	HighAction Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction `json:"highAction,omitempty" yaml:"highAction,omitempty"`

	// Action to take for a low risk. See action block below.
	LowAction Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction `json:"lowAction,omitempty" yaml:"lowAction,omitempty"`

	// Action to take for a medium risk. See action block below.
	MediumAction Cognito_RiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction `json:"mediumAction,omitempty" yaml:"mediumAction,omitempty"`
}
