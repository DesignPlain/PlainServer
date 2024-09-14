package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypes struct {
	/*
	   If a finding is matched by any of the infoType detectors listed here, the finding will be excluded from the scan results.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypesInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`
}
