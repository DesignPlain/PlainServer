package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleKeywordMatchConfiguration struct {
	// Collection of keywords to match.
	Keywords []string `json:"keywords,omitempty" yaml:"keywords,omitempty"`

	// Negate the rule.
	Negate bool `json:"negate,omitempty" yaml:"negate,omitempty"`

	// Rule name.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`
}
