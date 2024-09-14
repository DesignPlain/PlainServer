package types

type Container_AzureClusterLoggingConfig struct {
	// Configuration of the logging components.
	ComponentConfig Container_AzureClusterLoggingConfigComponentConfig `json:"componentConfig,omitempty" yaml:"componentConfig,omitempty"`
}
