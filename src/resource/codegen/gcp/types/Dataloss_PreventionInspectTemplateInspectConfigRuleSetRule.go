package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRule struct {
	/*
	   The rule that specifies conditions when findings of infoTypes specified in InspectionRuleSet are removed from results.
	   Structure is documented below.
	*/
	ExclusionRule Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRule `json:"exclusionRule,omitempty" yaml:"exclusionRule,omitempty"`

	/*
	   Hotword-based detection rule.
	   Structure is documented below.
	*/
	HotwordRule Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRule `json:"hotwordRule,omitempty" yaml:"hotwordRule,omitempty"`
}
