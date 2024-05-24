package types

type Cognito_RiskConfigurationCompromisedCredentialsRiskConfiguration struct {
	// The compromised credentials risk configuration actions. See details below.
	Actions Cognito_RiskConfigurationCompromisedCredentialsRiskConfigurationActions `json:"actions,omitempty" yaml:"actions,omitempty"`

	// Perform the action for these events. The default is to perform all events if no event filter is specified. Valid values are `SIGN_IN`, `PASSWORD_CHANGE`, and `SIGN_UP`.
	EventFilters []string `json:"eventFilters,omitempty" yaml:"eventFilters,omitempty"`
}
