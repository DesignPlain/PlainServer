package types

type Container_AttachedClusterLoggingConfig struct {
	/*
	   The configuration of the logging components
	   Structure is documented below.
	*/
	ComponentConfig Container_AttachedClusterLoggingConfigComponentConfig `json:"componentConfig,omitempty" yaml:"componentConfig,omitempty"`
}
