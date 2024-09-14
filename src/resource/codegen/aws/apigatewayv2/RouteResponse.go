package apigatewayv2

type RouteResponse struct {
	// Response models for the route response.
	ResponseModels map[string]string `json:"responseModels,omitempty" yaml:"responseModels,omitempty"`

	// Identifier of the `aws.apigatewayv2.Route`.
	RouteId string `json:"routeId,omitempty" yaml:"routeId,omitempty"`

	// Route response key.
	RouteResponseKey string `json:"routeResponseKey,omitempty" yaml:"routeResponseKey,omitempty"`

	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route response.
	ModelSelectionExpression string `json:"modelSelectionExpression,omitempty" yaml:"modelSelectionExpression,omitempty"`
}
