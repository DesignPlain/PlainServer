package apigateway

type RestApiPolicy struct {
	// JSON formatted policy document that controls access to the API Gateway.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// ID of the REST API.
	RestApiId string `json:"restApiId,omitempty" yaml:"restApiId,omitempty"`
}
