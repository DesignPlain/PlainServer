package types

type Apigateway_DocumentationPartLocation struct {
	// HTTP status code of a response. The default value is `-` for any status code.
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	// Type of API entity to which the documentation content appliesE.g., `API`, `METHOD` or `REQUEST_BODY`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// HTTP verb of a method. The default value is `-` for any method.
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	// Name of the targeted API entity.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// URL path of the target. The default value is `/` for the root resource.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
