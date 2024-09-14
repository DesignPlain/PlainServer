package apigatewayv2

type ApiMapping struct {
	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// The API mapping key. Refer to [REST API](https://docs.aws.amazon.com/apigateway/latest/developerguide/rest-api-mappings.html), [HTTP API](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-mappings.html) or [WebSocket API](https://docs.aws.amazon.com/apigateway/latest/developerguide/websocket-api-mappings.html).
	ApiMappingKey string `json:"apiMappingKey,omitempty" yaml:"apiMappingKey,omitempty"`

	// Domain name. Use the `aws.apigatewayv2.DomainName` resource to configure a domain name.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// API stage. Use the `aws.apigatewayv2.Stage` resource to configure an API stage.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`
}
