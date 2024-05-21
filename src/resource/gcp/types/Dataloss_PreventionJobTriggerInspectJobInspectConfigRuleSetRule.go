package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRule struct {
	/*
	   The rule that specifies conditions when findings of infoTypes specified in InspectionRuleSet are removed from results.
	   Structure is documented below.
	*/
	ExclusionRule Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRule `json:"exclusionRule,omitempty" yaml:"exclusionRule,omitempty"`

	/*
	   Hotword-based detection rule.
	   Structure is documented below.
	*/
	HotwordRule Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleHotwordRule `json:"hotwordRule,omitempty" yaml:"hotwordRule,omitempty"`
}
