package types

type Dataloss_PreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType struct {
	/*
	   Type of information the findings limit applies to. Only one limit per infoType should be provided. If InfoTypeLimit does
	   not have an infoType, the DLP API applies the limit against all infoTypes that are found but not
	   specified in another InfoTypeLimit.
	   Structure is documented below.
	*/
	InfoType Dataloss_PreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `json:"infoType,omitempty" yaml:"infoType,omitempty"`

	// Max findings limit for the given infoType.
	MaxFindings int `json:"maxFindings,omitempty" yaml:"maxFindings,omitempty"`
}
