package types

type Wafv2_WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition struct {
	// Action setting that a log record must contain in order to meet the condition. Valid values for `action` are `ALLOW`, `BLOCK`, and `COUNT`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
}
