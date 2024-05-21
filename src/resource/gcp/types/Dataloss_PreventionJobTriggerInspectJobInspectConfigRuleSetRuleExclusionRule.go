package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRule struct {
	/*
	   Dictionary which defines the rule.
	   Structure is documented below.
	*/
	Dictionary Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionary `json:"dictionary,omitempty" yaml:"dictionary,omitempty"`

	/*
	   Drop if the hotword rule is contained in the proximate context.
	   Structure is documented below.
	*/
	ExcludeByHotword Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleExcludeByHotword `json:"excludeByHotword,omitempty" yaml:"excludeByHotword,omitempty"`

	/*
	   Set of infoTypes for which findings would affect this rule.
	   Structure is documented below.
	*/
	ExcludeInfoTypes Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypes `json:"excludeInfoTypes,omitempty" yaml:"excludeInfoTypes,omitempty"`

	/*
	   How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType
	   Possible values are: `MATCHING_TYPE_FULL_MATCH`, `MATCHING_TYPE_PARTIAL_MATCH`, `MATCHING_TYPE_INVERSE_MATCH`.
	*/
	MatchingType string `json:"matchingType,omitempty" yaml:"matchingType,omitempty"`

	/*
	   Regular expression which defines the rule.
	   Structure is documented below.
	*/
	Regex Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleRegex `json:"regex,omitempty" yaml:"regex,omitempty"`
}
