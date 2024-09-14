package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRule struct {
	// Configuration for a sentiment rule.
	SentimentConfiguration Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleSentimentConfiguration `json:"sentimentConfiguration,omitempty" yaml:"sentimentConfiguration,omitempty"`

	// Rule type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Configuration for an issue detection rule.
	IssueDetectionConfiguration Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleIssueDetectionConfiguration `json:"issueDetectionConfiguration,omitempty" yaml:"issueDetectionConfiguration,omitempty"`

	// Configuration for a keyword match rule.
	KeywordMatchConfiguration Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleKeywordMatchConfiguration `json:"keywordMatchConfiguration,omitempty" yaml:"keywordMatchConfiguration,omitempty"`
}
