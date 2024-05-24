package types

type Ecs_ServiceServiceConnectConfiguration struct {
	// The log configuration for the container. See below.
	LogConfiguration Ecs_ServiceServiceConnectConfigurationLogConfiguration `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`

	// The namespace name or ARN of the `aws.servicediscovery.HttpNamespace` for use with Service Connect.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// The list of Service Connect service objects. See below.
	Services []Ecs_ServiceServiceConnectConfigurationService `json:"services,omitempty" yaml:"services,omitempty"`

	// Specifies whether to use Service Connect with this service.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
