package types

type Apigatewayv2_IntegrationResponseParameter struct {
	/*
	   Key-value map. The key of this map identifies the location of the request parameter to change, and how to change it. The corresponding value specifies the new data for the parameter.
	   See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	*/
	Mappings map[string]string `json:"mappings,omitempty" yaml:"mappings,omitempty"`

	// HTTP status code in the range 200-599.
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
