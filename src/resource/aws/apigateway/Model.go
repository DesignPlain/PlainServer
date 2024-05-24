package apigateway

type Model struct {
	// Content type of the model
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Description of the model
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the model
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ID of the associated REST API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Schema of the model in a JSON form
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
