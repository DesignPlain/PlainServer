package types

type Apprunner_ServiceObservabilityConfiguration struct {
	// ARN of the observability configuration that is associated with the service. Specified only when `observability_enabled` is `true`.
	ObservabilityConfigurationArn string `json:"observabilityConfigurationArn,omitempty" yaml:"observabilityConfigurationArn,omitempty"`

	// When `true`, an observability configuration resource is associated with the service.
	ObservabilityEnabled bool `json:"observabilityEnabled,omitempty" yaml:"observabilityEnabled,omitempty"`
}
