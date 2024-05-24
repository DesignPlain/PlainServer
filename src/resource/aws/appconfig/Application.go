package appconfig

type Application struct {
	// Description of the application. Can be at most 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name for the application. Must be between 1 and 64 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
