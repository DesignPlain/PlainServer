package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypes struct {
	/*
	   If a finding is matched by any of the infoType detectors listed here, the finding will be excluded from the scan results.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypesInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`
}
