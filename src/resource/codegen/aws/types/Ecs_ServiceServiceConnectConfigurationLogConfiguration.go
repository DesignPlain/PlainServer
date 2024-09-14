package types

type Ecs_ServiceServiceConnectConfigurationLogConfiguration struct {
	// Log driver to use for the container.
	LogDriver string `json:"logDriver,omitempty" yaml:"logDriver,omitempty"`

	// Configuration options to send to the log driver.
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`

	// Secrets to pass to the log configuration. See below.
	SecretOptions []Ecs_ServiceServiceConnectConfigurationLogConfigurationSecretOption `json:"secretOptions,omitempty" yaml:"secretOptions,omitempty"`
}
