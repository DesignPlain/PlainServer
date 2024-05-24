package apigatewayv2

import types "DesignSphere_Server/src/resource/aws/types"

type Api struct {
	// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn string `json:"credentialsArn,omitempty" yaml:"credentialsArn,omitempty"`

	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
	FailOnWarnings bool `json:"failOnWarnings,omitempty" yaml:"failOnWarnings,omitempty"`

	// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
	RouteKey string `json:"routeKey,omitempty" yaml:"routeKey,omitempty"`

	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	   Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
	   Applicable for WebSocket APIs.
	*/
	ApiKeySelectionExpression string `json:"apiKeySelectionExpression,omitempty" yaml:"apiKeySelectionExpression,omitempty"`

	/*
	   Whether clients can invoke the API by using the default `execute-api` endpoint.
	   By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
	   To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	*/
	DisableExecuteApiEndpoint bool `json:"disableExecuteApiEndpoint,omitempty" yaml:"disableExecuteApiEndpoint,omitempty"`

	// Name of the API. Must be less than or equal to 128 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
	ProtocolType string `json:"protocolType,omitempty" yaml:"protocolType,omitempty"`

	/*
	   The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
	   Defaults to `$request.method $request.path`.
	*/
	RouteSelectionExpression string `json:"routeSelectionExpression,omitempty" yaml:"routeSelectionExpression,omitempty"`

	// Map of tags to assign to the API. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
	CorsConfiguration types.Apigatewayv2_ApiCorsConfiguration `json:"corsConfiguration,omitempty" yaml:"corsConfiguration,omitempty"`

	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	   For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	   The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
	*/
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body string `json:"body,omitempty" yaml:"body,omitempty"`
}
