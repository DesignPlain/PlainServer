package types

type Lambda_getFunctionLoggingConfig struct {
	//
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`

	//
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`

	//
	SystemLogLevel string `json:"systemLogLevel,omitempty" yaml:"systemLogLevel,omitempty"`

	//
	ApplicationLogLevel string `json:"applicationLogLevel,omitempty" yaml:"applicationLogLevel,omitempty"`
}
