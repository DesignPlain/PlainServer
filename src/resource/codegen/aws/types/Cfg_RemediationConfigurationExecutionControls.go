package types

type Cfg_RemediationConfigurationExecutionControls struct {
	// Configuration block for SSM controls. See below.
	SsmControls Cfg_RemediationConfigurationExecutionControlsSsmControls `json:"ssmControls,omitempty" yaml:"ssmControls,omitempty"`
}
