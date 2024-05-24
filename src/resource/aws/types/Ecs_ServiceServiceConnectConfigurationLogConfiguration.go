package types

type Ecs_ServiceServiceConnectConfigurationLogConfiguration struct {
	// The log driver to use for the container.
	LogDriver string `json:"logDriver,omitempty" yaml:"logDriver,omitempty"`

	// The configuration options to send to the log driver.
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`

	// The secrets to pass to the log configuration. See below.
	SecretOptions []Ecs_ServiceServiceConnectConfigurationLogConfigurationSecretOption `json:"secretOptions,omitempty" yaml:"secretOptions,omitempty"`
}
