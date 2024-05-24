package finspace

type KxDatabase struct {
	// Description of the KX database.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unique identifier for the KX environment.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	/*
	   Name of the KX database.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
