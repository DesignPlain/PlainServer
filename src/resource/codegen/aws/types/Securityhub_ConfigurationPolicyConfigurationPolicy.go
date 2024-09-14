package types

type Securityhub_ConfigurationPolicyConfigurationPolicy struct {
	// A list that defines which security standards are enabled in the configuration policy. It must be defined if `service_enabled` is set to true.
	EnabledStandardArns []string `json:"enabledStandardArns,omitempty" yaml:"enabledStandardArns,omitempty"`

	// Defines which security controls are enabled in the configuration policy and any customizations to parameters affecting them. See below.
	SecurityControlsConfiguration Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfiguration `json:"securityControlsConfiguration,omitempty" yaml:"securityControlsConfiguration,omitempty"`

	// Indicates whether Security Hub is enabled in the policy.
	ServiceEnabled bool `json:"serviceEnabled,omitempty" yaml:"serviceEnabled,omitempty"`
}
