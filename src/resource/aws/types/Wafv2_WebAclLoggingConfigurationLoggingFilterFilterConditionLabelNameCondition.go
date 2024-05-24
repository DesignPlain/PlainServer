package types

type Wafv2_WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition struct {
	// Name of the label that a log record must contain in order to meet the condition. It must be a fully qualified label name, which includes a prefix, optional namespaces, and the label name itself. The prefix identifies the rule group or web ACL context of the rule that added the label.
	LabelName string `json:"labelName,omitempty" yaml:"labelName,omitempty"`
}
