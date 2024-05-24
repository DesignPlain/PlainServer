package apigatewayv2

import types "DesignSphere_Server/src/resource/aws/types"

type Authorizer struct {
	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	/*
	   Required credentials as an IAM role for API Gateway to invoke the authorizer.
	   Supported only for `REQUEST` authorizers.
	*/
	AuthorizerCredentialsArn string `json:"authorizerCredentialsArn,omitempty" yaml:"authorizerCredentialsArn,omitempty"`

	/*
	   Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
	   Supported only for HTTP APIs.
	*/
	JwtConfiguration types.Apigatewayv2_AuthorizerJwtConfiguration `json:"jwtConfiguration,omitempty" yaml:"jwtConfiguration,omitempty"`

	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	   Valid values: `1.0`, `2.0`.
	*/
	AuthorizerPayloadFormatVersion string `json:"authorizerPayloadFormatVersion,omitempty" yaml:"authorizerPayloadFormatVersion,omitempty"`

	/*
	   Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	   If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
	   Supported only for HTTP API Lambda authorizers.
	*/
	AuthorizerResultTtlInSeconds int `json:"authorizerResultTtlInSeconds,omitempty" yaml:"authorizerResultTtlInSeconds,omitempty"`

	/*
	   Authorizer type. Valid values: `JWT`, `REQUEST`.
	   Specify `REQUEST` for a Lambda function using incoming request parameters.
	   For HTTP APIs, specify `JWT` to use JSON Web Tokens.
	*/
	AuthorizerType string `json:"authorizerType,omitempty" yaml:"authorizerType,omitempty"`

	/*
	   Authorizer's Uniform Resource Identifier (URI).
	   For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `aws.lambda.Function` resource.
	   Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
	*/
	AuthorizerUri string `json:"authorizerUri,omitempty" yaml:"authorizerUri,omitempty"`

	/*
	   Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	   Supported only for HTTP APIs.
	*/
	EnableSimpleResponses bool `json:"enableSimpleResponses,omitempty" yaml:"enableSimpleResponses,omitempty"`

	/*
	   Identity sources for which authorization is requested.
	   For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	   For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	*/
	IdentitySources []string `json:"identitySources,omitempty" yaml:"identitySources,omitempty"`
}
