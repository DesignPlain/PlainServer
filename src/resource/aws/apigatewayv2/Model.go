package apigatewayv2

type Model struct {
	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// The content-type for the model, for example, `application/json`. Must be between 1 and 256 characters in length.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Description of the model. Must be between 1 and 128 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the model. Must be alphanumeric. Must be between 1 and 128 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Schema for the model. This should be a [JSON schema draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Must be less than or equal to 32768 characters in length.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
