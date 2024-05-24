package apprunner

type Connection struct {
	// Name of the connection.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// Source repository provider. Valid values: `GITHUB`.
	ProviderType string `json:"providerType,omitempty" yaml:"providerType,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
