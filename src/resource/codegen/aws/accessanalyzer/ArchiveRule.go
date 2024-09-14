package accessanalyzer

import types "libds/aws/types"

type ArchiveRule struct {
	// Rule name.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	// Analyzer name.
	AnalyzerName string `json:"analyzerName,omitempty" yaml:"analyzerName,omitempty"`

	// Filter criteria for the archive rule. See Filter for more details.
	Filters []types.Accessanalyzer_ArchiveRuleFilter `json:"filters,omitempty" yaml:"filters,omitempty"`
}
