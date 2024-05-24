package apigatewayv2

type IntegrationResponse struct {
	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
	ContentHandlingStrategy string `json:"contentHandlingStrategy,omitempty" yaml:"contentHandlingStrategy,omitempty"`

	// Identifier of the `aws.apigatewayv2.Integration`.
	IntegrationId string `json:"integrationId,omitempty" yaml:"integrationId,omitempty"`

	// Integration response key.
	IntegrationResponseKey string `json:"integrationResponseKey,omitempty" yaml:"integrationResponseKey,omitempty"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" yaml:"responseTemplates,omitempty"`

	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
	TemplateSelectionExpression string `json:"templateSelectionExpression,omitempty" yaml:"templateSelectionExpression,omitempty"`
}
