package servicediscovery

type PublicDnsNamespace struct {
	// The description that you specify for the namespace when you create it.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the namespace.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the namespace. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
