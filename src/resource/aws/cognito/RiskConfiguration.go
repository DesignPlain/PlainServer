package cognito

import types "DesignSphere_Server/src/resource/aws/types"

type RiskConfiguration struct {
	// The account takeover risk configuration. See details below.
	AccountTakeoverRiskConfiguration types.Cognito_RiskConfigurationAccountTakeoverRiskConfiguration `json:"accountTakeoverRiskConfiguration,omitempty" yaml:"accountTakeoverRiskConfiguration,omitempty"`

	// The app client ID. When the client ID is not provided, the same risk configuration is applied to all the clients in the User Pool.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The compromised credentials risk configuration. See details below.
	CompromisedCredentialsRiskConfiguration types.Cognito_RiskConfigurationCompromisedCredentialsRiskConfiguration `json:"compromisedCredentialsRiskConfiguration,omitempty" yaml:"compromisedCredentialsRiskConfiguration,omitempty"`

	// The configuration to override the risk decision. See details below.
	RiskExceptionConfiguration types.Cognito_RiskConfigurationRiskExceptionConfiguration `json:"riskExceptionConfiguration,omitempty" yaml:"riskExceptionConfiguration,omitempty"`

	// The user pool ID.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
