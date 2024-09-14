package schemas

type Schema struct {
	// The name of the registry in which this schema belongs.
	RegistryName string `json:"registryName,omitempty" yaml:"registryName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The schema specification. Must be a valid Open API 3.0 spec.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// The description of the schema. Maximum of 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
