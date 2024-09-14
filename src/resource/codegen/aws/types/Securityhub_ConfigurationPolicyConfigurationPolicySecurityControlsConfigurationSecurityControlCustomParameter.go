package types

type Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameter struct {
	// An object that specifies parameter values for a control in a configuration policy. See below.
	Parameters []Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The ID of the security control. For more information see the [Security Hub controls reference] documentation.
	SecurityControlId string `json:"securityControlId,omitempty" yaml:"securityControlId,omitempty"`
}
