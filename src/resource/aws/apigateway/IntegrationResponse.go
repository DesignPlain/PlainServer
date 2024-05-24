package apigateway

type IntegrationResponse struct {
	// ID of the associated REST API.
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Regular expression pattern used to choose an integration response based on the response from the backend. Omit configuring this to make the integration the default one. If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched. For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
	SelectionPattern string `json:"selectionPattern,omitempty" yaml:"selectionPattern,omitempty"`

	/*
	   HTTP status code.

	   The following arguments are optional:
	*/
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	// How to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
	ContentHandling string `json:"contentHandling,omitempty" yaml:"contentHandling,omitempty"`

	// HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`).
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	// API resource ID.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Map of response parameters that can be read from the backend response. For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`.
	ResponseParameters map[string]string `json:"responseParameters,omitempty" yaml:"responseParameters,omitempty"`

	// Map of templates used to transform the integration response body.
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" yaml:"responseTemplates,omitempty"`
}
