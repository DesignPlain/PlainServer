package ivschat

import types "DesignSphere_Server/src/resource/aws/types"

type LoggingConfiguration struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Object containing destination configuration for where chat activity will be logged. This object must contain exactly one of the following children arguments:
	DestinationConfiguration types.Ivschat_LoggingConfigurationDestinationConfiguration `json:"destinationConfiguration,omitempty" yaml:"destinationConfiguration,omitempty"`

	// Logging Configuration name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
