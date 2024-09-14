package apigateway

type Authorizer struct {
	/*
	   Authorizer's Uniform Resource Identifier (URI). This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	   e.g., `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	*/
	AuthorizerUri string `json:"authorizerUri,omitempty" yaml:"authorizerUri,omitempty"`

	// Source of the identity in an incoming request. Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g., `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource string `json:"identitySource,omitempty" yaml:"identitySource,omitempty"`

	// Validation expression for the incoming identity. For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched against this expression, and will proceed if the token matches. If the token doesn't match, the client receives a 401 Unauthorized response.
	IdentityValidationExpression string `json:"identityValidationExpression,omitempty" yaml:"identityValidationExpression,omitempty"`

	// Name of the authorizer
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ID of the associated REST API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool. Defaults to `TOKEN`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Credentials required for the authorizer. To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials string `json:"authorizerCredentials,omitempty" yaml:"authorizerCredentials,omitempty"`

	// TTL of cached authorizer results in seconds. Defaults to `300`.
	AuthorizerResultTtlInSeconds int `json:"authorizerResultTtlInSeconds,omitempty" yaml:"authorizerResultTtlInSeconds,omitempty"`

	// List of the Amazon Cognito user pool ARNs. Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns []string `json:"providerArns,omitempty" yaml:"providerArns,omitempty"`
}
