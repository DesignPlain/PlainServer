package apigatewayv2

import types "libds/aws/types"

type Route struct {
	// Boolean whether an API key is required for the route. Defaults to `false`. Supported only for WebSocket APIs.
	ApiKeyRequired bool `json:"apiKeyRequired,omitempty" yaml:"apiKeyRequired,omitempty"`

	// Request models for the route. Supported only for WebSocket APIs.
	RequestModels map[string]string `json:"requestModels,omitempty" yaml:"requestModels,omitempty"`

	// Route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
	RouteKey string `json:"routeKey,omitempty" yaml:"routeKey,omitempty"`

	// The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route. Supported only for WebSocket APIs.
	RouteResponseSelectionExpression string `json:"routeResponseSelectionExpression,omitempty" yaml:"routeResponseSelectionExpression,omitempty"`

	// Operation name for the route. Must be between 1 and 64 characters in length.
	OperationName string `json:"operationName,omitempty" yaml:"operationName,omitempty"`

	// Request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters []types.Apigatewayv2_RouteRequestParameter `json:"requestParameters,omitempty" yaml:"requestParameters,omitempty"`

	// Target for the route, of the form `integrations/`-`IntegrationID`-, where -`IntegrationID`- is the identifier of an `aws.apigatewayv2.Integration` resource.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// Authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	AuthorizationScopes []string `json:"authorizationScopes,omitempty" yaml:"authorizationScopes,omitempty"`

	/*
	   Authorization type for the route.
	   For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	   For HTTP APIs, valid values are `NONE` for open access, `JWT` for using JSON Web Tokens, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	   Defaults to `NONE`.
	*/
	AuthorizationType string `json:"authorizationType,omitempty" yaml:"authorizationType,omitempty"`

	// Identifier of the `aws.apigatewayv2.Authorizer` resource to be associated with this route.
	AuthorizerId string `json:"authorizerId,omitempty" yaml:"authorizerId,omitempty"`

	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression string `json:"modelSelectionExpression,omitempty" yaml:"modelSelectionExpression,omitempty"`
}
