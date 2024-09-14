package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRule struct {
	/*
	   How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType
	   Possible values are: `MATCHING_TYPE_FULL_MATCH`, `MATCHING_TYPE_PARTIAL_MATCH`, `MATCHING_TYPE_INVERSE_MATCH`.
	*/
	MatchingType string `json:"matchingType,omitempty" yaml:"matchingType,omitempty"`

	/*
	   Regular expression which defines the rule.
	   Structure is documented below.
	*/
	Regex Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleRegex `json:"regex,omitempty" yaml:"regex,omitempty"`

	/*
	   Dictionary which defines the rule.
	   Structure is documented below.
	*/
	Dictionary Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleDictionary `json:"dictionary,omitempty" yaml:"dictionary,omitempty"`

	/*
	   Drop if the hotword rule is contained in the proximate context.
	   For tabular data, the context includes the column name.
	   Structure is documented below.
	*/
	ExcludeByHotword Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotword `json:"excludeByHotword,omitempty" yaml:"excludeByHotword,omitempty"`

	/*
	   Set of infoTypes for which findings would affect this rule.
	   Structure is documented below.
	*/
	ExcludeInfoTypes Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypes `json:"excludeInfoTypes,omitempty" yaml:"excludeInfoTypes,omitempty"`
}
