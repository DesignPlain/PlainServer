package types

type Iot_DomainConfigurationAuthorizerConfig struct {
	// A Boolean that specifies whether the domain configuration's authorization service can be overridden.
	AllowAuthorizerOverride bool `json:"allowAuthorizerOverride,omitempty" yaml:"allowAuthorizerOverride,omitempty"`

	// The name of the authorization service for a domain configuration.
	DefaultAuthorizerName string `json:"defaultAuthorizerName,omitempty" yaml:"defaultAuthorizerName,omitempty"`
}
