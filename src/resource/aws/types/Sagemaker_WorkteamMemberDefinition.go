package types

type Sagemaker_WorkteamMemberDefinition struct {
	// A list user groups that exist in your OIDC Identity Provider (IdP). One to ten groups can be used to create a single private work team. See Cognito Member Definition details below.
	OidcMemberDefinition Sagemaker_WorkteamMemberDefinitionOidcMemberDefinition `json:"oidcMemberDefinition,omitempty" yaml:"oidcMemberDefinition,omitempty"`

	// The Amazon Cognito user group that is part of the work team. See Cognito Member Definition details below.
	CognitoMemberDefinition Sagemaker_WorkteamMemberDefinitionCognitoMemberDefinition `json:"cognitoMemberDefinition,omitempty" yaml:"cognitoMemberDefinition,omitempty"`
}
