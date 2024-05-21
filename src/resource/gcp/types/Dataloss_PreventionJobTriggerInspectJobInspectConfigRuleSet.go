package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSet struct {
	/*
	   List of infoTypes this rule set is applied to.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`

	/*
	   Set of rules to be applied to info The rules are applied in order.
	   Structure is documented below.
	*/
	Rules []Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
