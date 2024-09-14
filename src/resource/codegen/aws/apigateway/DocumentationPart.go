package apigateway

import types "libds/aws/types"

type DocumentationPart struct {
	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location types.Apigateway_DocumentationPartLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// ID of the associated Rest API
	RestApiId string `json:"restApiId,omitempty" yaml:"restApiId,omitempty"`
}
