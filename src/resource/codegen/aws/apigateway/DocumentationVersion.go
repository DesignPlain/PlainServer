package apigateway

type DocumentationVersion struct {
	// Description of the API documentation version.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ID of the associated Rest API
	RestApiId string `json:"restApiId,omitempty" yaml:"restApiId,omitempty"`

	// Version identifier of the API documentation snapshot.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
