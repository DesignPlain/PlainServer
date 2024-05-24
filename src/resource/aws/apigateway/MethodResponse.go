package apigateway

type MethodResponse struct {
	/*
	   A map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header name and the associated value is a boolean flag indicating whether the method response parameter is required. The method response header names must match the pattern of `method.response.header.{name}`, where `name` is a valid and unique header name.

	   The response parameter names defined here are available in the integration response to be mapped from an integration response header expressed in `integration.response.header.{name}`, a static value enclosed within a pair of single quotes (e.g., '`application/json'`), or a JSON expression from the back-end response payload in the form of `integration.response.body.{JSON-expression}`, where `JSON-expression` is a valid JSON expression without the `$` prefix.)
	*/
	ResponseParameters map[string]bool `json:"responseParameters,omitempty" yaml:"responseParameters,omitempty"`

	// The string identifier of the associated REST API.
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// The method response's status code.
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	// The HTTP verb of the method resource (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`).
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	// The Resource identifier for the method resource.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// A map specifying the model resources used for the response's content type. Response models are represented as a key/value map, with a content type as the key and a Model name as the value.
	ResponseModels map[string]string `json:"responseModels,omitempty" yaml:"responseModels,omitempty"`
}
