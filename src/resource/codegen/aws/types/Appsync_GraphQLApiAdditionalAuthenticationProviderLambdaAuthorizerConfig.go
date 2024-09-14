package types

type Appsync_GraphQLApiAdditionalAuthenticationProviderLambdaAuthorizerConfig struct {
	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a `ttlOverride` key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	AuthorizerResultTtlInSeconds int `json:"authorizerResultTtlInSeconds,omitempty" yaml:"authorizerResultTtlInSeconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow `lambda:InvokeFunction` from service principal `appsync.amazonaws.com`.
	AuthorizerUri string `json:"authorizerUri,omitempty" yaml:"authorizerUri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression string `json:"identityValidationExpression,omitempty" yaml:"identityValidationExpression,omitempty"`
}
