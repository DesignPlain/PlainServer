package types

type Wafv2_WebAclLoggingConfigurationLoggingFilterFilterCondition struct {
	// Configuration for a single action condition. See Action Condition below for more details.
	ActionCondition Wafv2_WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition `json:"actionCondition,omitempty" yaml:"actionCondition,omitempty"`

	// Condition for a single label name. See Label Name Condition below for more details.
	LabelNameCondition Wafv2_WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition `json:"labelNameCondition,omitempty" yaml:"labelNameCondition,omitempty"`
}
