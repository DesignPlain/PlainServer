package types

type Appsync_GraphQLApiUserPoolConfig struct {
	// AWS region in which the user pool was created.
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// Action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
	DefaultAction string `json:"defaultAction,omitempty" yaml:"defaultAction,omitempty"`

	// User pool ID.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	AppIdClientRegex string `json:"appIdClientRegex,omitempty" yaml:"appIdClientRegex,omitempty"`
}
