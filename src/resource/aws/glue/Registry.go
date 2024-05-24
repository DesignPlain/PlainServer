package glue

type Registry struct {
	// A description of the registry.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The Name of the registry.
	RegistryName string `json:"registryName,omitempty" yaml:"registryName,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
