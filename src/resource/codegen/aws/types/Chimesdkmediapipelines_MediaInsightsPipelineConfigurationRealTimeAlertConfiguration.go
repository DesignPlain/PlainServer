package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfiguration struct {
	// Disables real time alert rules.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Collection of real time alert rules
	Rules []Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
