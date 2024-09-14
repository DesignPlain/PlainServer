package types

type Appsync_GraphQLApiAdditionalAuthenticationProviderUserPoolConfig struct {
	// AWS region in which the user pool was created.
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// User pool ID.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	AppIdClientRegex string `json:"appIdClientRegex,omitempty" yaml:"appIdClientRegex,omitempty"`
}
