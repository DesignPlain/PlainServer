package types

type Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfiguration struct {
	// A list of security controls that are disabled in the configuration policy Security Hub enables all other controls (including newly released controls) other than the listed controls. Conflicts with `enabled_control_identifiers`.
	DisabledControlIdentifiers []string `json:"disabledControlIdentifiers,omitempty" yaml:"disabledControlIdentifiers,omitempty"`

	// A list of security controls that are enabled in the configuration policy. Security Hub disables all other controls (including newly released controls) other than the listed controls. Conflicts with `disabled_control_identifiers`.
	EnabledControlIdentifiers []string `json:"enabledControlIdentifiers,omitempty" yaml:"enabledControlIdentifiers,omitempty"`

	// A list of control parameter customizations that are included in a configuration policy. Include multiple blocks to define multiple control custom parameters. See below.
	SecurityControlCustomParameters []Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameter `json:"securityControlCustomParameters,omitempty" yaml:"securityControlCustomParameters,omitempty"`
}
