package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLogConfiguration struct {
	// The log driver to use for the container.
	LogDriver string `json:"logDriver,omitempty" yaml:"logDriver,omitempty"`

	// The configuration options to send to the log driver.
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`

	// The secrets to pass to the log configuration.
	SecretOptions []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLogConfigurationSecretOption `json:"secretOptions,omitempty" yaml:"secretOptions,omitempty"`
}
