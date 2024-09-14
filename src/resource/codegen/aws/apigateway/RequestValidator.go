package apigateway

type RequestValidator struct {
	// Name of the request validator
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ID of the associated Rest API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Boolean whether to validate request body. Defaults to `false`.
	ValidateRequestBody bool `json:"validateRequestBody,omitempty" yaml:"validateRequestBody,omitempty"`

	// Boolean whether to validate request parameters. Defaults to `false`.
	ValidateRequestParameters bool `json:"validateRequestParameters,omitempty" yaml:"validateRequestParameters,omitempty"`
}
