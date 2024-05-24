package types

type Glue_UserDefinedFunctionResourceUri struct {
	// The type of the resource. can be one of `JAR`, `FILE`, and `ARCHIVE`.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// The URI for accessing the resource.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
