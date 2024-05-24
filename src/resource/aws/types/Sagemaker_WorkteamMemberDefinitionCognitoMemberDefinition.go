package types

type Sagemaker_WorkteamMemberDefinitionCognitoMemberDefinition struct {
	// An identifier for an application client. You must create the app client ID using Amazon Cognito.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// An identifier for a user group.
	UserGroup string `json:"userGroup,omitempty" yaml:"userGroup,omitempty"`

	// An identifier for a user pool. The user pool must be in the same region as the service that you are calling.
	UserPool string `json:"userPool,omitempty" yaml:"userPool,omitempty"`
}
