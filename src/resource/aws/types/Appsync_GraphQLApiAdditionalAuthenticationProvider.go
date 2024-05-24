package types

type Appsync_GraphQLApiAdditionalAuthenticationProvider struct {
	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig Appsync_GraphQLApiAdditionalAuthenticationProviderLambdaAuthorizerConfig `json:"lambdaAuthorizerConfig,omitempty" yaml:"lambdaAuthorizerConfig,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig Appsync_GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig `json:"openidConnectConfig,omitempty" yaml:"openidConnectConfig,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig Appsync_GraphQLApiAdditionalAuthenticationProviderUserPoolConfig `json:"userPoolConfig,omitempty" yaml:"userPoolConfig,omitempty"`
}
