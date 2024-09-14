package types

type Accessanalyzer_AnalyzerConfiguration struct {
	// A block that specifies the configuration of an unused access analyzer for an AWS organization or account. Documented below
	UnusedAccess Accessanalyzer_AnalyzerConfigurationUnusedAccess `json:"unusedAccess,omitempty" yaml:"unusedAccess,omitempty"`
}
