package apigateway

type Response struct {
	// Map of parameters (paths, query strings and headers) of the Gateway Response.
	ResponseParameters map[string]string `json:"responseParameters,omitempty" yaml:"responseParameters,omitempty"`

	// Map of templates used to transform the response body.
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" yaml:"responseTemplates,omitempty"`

	// Response type of the associated GatewayResponse.
	ResponseType string `json:"responseType,omitempty" yaml:"responseType,omitempty"`

	// String identifier of the associated REST API.
	RestApiId string `json:"restApiId,omitempty" yaml:"restApiId,omitempty"`

	// HTTP status code of the Gateway Response.
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
