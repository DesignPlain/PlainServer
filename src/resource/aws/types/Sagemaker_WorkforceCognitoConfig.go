package types

type Sagemaker_WorkforceCognitoConfig struct {
	// The client ID for your Amazon Cognito user pool.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// ID for your Amazon Cognito user pool.
	UserPool string `json:"userPool,omitempty" yaml:"userPool,omitempty"`
}
