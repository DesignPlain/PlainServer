package accessanalyzer

import types "DesignSphere_Server/src/resource/aws/types"

type ArchiveRule struct {
	// Analyzer name.
	AnalyzerName string `json:"analyzerName,omitempty" yaml:"analyzerName,omitempty"`

	// Filter criteria for the archive rule. See Filter for more details.
	Filters []types.Accessanalyzer_ArchiveRuleFilter `json:"filters,omitempty" yaml:"filters,omitempty"`

	// Rule name.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`
}
