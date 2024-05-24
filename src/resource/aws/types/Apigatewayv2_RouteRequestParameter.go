package types

type Apigatewayv2_RouteRequestParameter struct {
	// Request parameter key. This is a [request data mapping parameter](https://docs.aws.amazon.com/apigateway/latest/developerguide/websocket-api-data-mapping.html#websocket-mapping-request-parameters).
	RequestParameterKey string `json:"requestParameterKey,omitempty" yaml:"requestParameterKey,omitempty"`

	// Boolean whether or not the parameter is required.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}
