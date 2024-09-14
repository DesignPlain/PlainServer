package bedrockmodel

import types "libds/aws/types"

type InvocationLoggingConfiguration struct {
	// The logging configuration values to set.
	LoggingConfig types.Bedrockmodel_InvocationLoggingConfigurationLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`
}
