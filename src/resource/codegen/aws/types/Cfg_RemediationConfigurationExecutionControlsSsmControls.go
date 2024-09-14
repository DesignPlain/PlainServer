package types

type Cfg_RemediationConfigurationExecutionControlsSsmControls struct {
	// Maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule. The default value is 10%!.(MISSING)
	ConcurrentExecutionRatePercentage int `json:"concurrentExecutionRatePercentage,omitempty" yaml:"concurrentExecutionRatePercentage,omitempty"`

	// Percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule. The default is 50%!.(MISSING)
	ErrorPercentage int `json:"errorPercentage,omitempty" yaml:"errorPercentage,omitempty"`
}
