package apigatewayv2

import types "DesignSphere_Server/src/resource/aws/types"

type Integration struct {
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy string `json:"contentHandlingStrategy,omitempty" yaml:"contentHandlingStrategy,omitempty"`

	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion string `json:"payloadFormatVersion,omitempty" yaml:"payloadFormatVersion,omitempty"`

	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression string `json:"templateSelectionExpression,omitempty" yaml:"templateSelectionExpression,omitempty"`

	// ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// Description of the integration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	   For HTTP APIs with a specified `integration_subtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
	   For HTTP APIs without a specified `integration_subtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	   See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	*/
	RequestParameters map[string]string `json:"requestParameters,omitempty" yaml:"requestParameters,omitempty"`

	// TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig types.Apigatewayv2_IntegrationTlsConfig `json:"tlsConfig,omitempty" yaml:"tlsConfig,omitempty"`

	// Integration's HTTP method. Must be specified if `integration_type` is not `MOCK`.
	IntegrationMethod string `json:"integrationMethod,omitempty" yaml:"integrationMethod,omitempty"`

	// AWS service action to invoke. Supported only for HTTP APIs when `integration_type` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype string `json:"integrationSubtype,omitempty" yaml:"integrationSubtype,omitempty"`

	/*
	   Integration type of an integration.
	   Valid values: `AWS` (supported only for WebSocket APIs), `AWS_PROXY`, `HTTP` (supported only for WebSocket APIs), `HTTP_PROXY`, `MOCK` (supported only for WebSocket APIs). For an HTTP API private integration, use `HTTP_PROXY`.
	*/
	IntegrationType string `json:"integrationType,omitempty" yaml:"integrationType,omitempty"`

	/*
	   URI of the Lambda function for a Lambda proxy integration, when `integration_type` is `AWS_PROXY`.
	   For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	*/
	IntegrationUri string `json:"integrationUri,omitempty" yaml:"integrationUri,omitempty"`

	// Map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]string `json:"requestTemplates,omitempty" yaml:"requestTemplates,omitempty"`

	/*
	   Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	   The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	   this provider will only perform drift detection of its value when present in a configuration.
	*/
	TimeoutMilliseconds int `json:"timeoutMilliseconds,omitempty" yaml:"timeoutMilliseconds,omitempty"`

	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// Type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType string `json:"connectionType,omitempty" yaml:"connectionType,omitempty"`

	// Credentials required for the integration, if any.
	CredentialsArn string `json:"credentialsArn,omitempty" yaml:"credentialsArn,omitempty"`

	/*
	   Pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `request_templates` attribute.
	   Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	*/
	PassthroughBehavior string `json:"passthroughBehavior,omitempty" yaml:"passthroughBehavior,omitempty"`

	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters []types.Apigatewayv2_IntegrationResponseParameter `json:"responseParameters,omitempty" yaml:"responseParameters,omitempty"`
}
