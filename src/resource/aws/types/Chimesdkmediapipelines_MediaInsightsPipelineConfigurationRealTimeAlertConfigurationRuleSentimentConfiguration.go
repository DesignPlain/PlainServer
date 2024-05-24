package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleSentimentConfiguration struct {
	// Rule name.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	// Sentiment type to match.
	SentimentType string `json:"sentimentType,omitempty" yaml:"sentimentType,omitempty"`

	// Analysis interval.
	TimePeriod int `json:"timePeriod,omitempty" yaml:"timePeriod,omitempty"`
}
