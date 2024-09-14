package types

type Dataloss_PreventionInspectTemplateInspectConfigCustomInfoTypeSensitivityScore struct {
	/*
	   The sensitivity score applied to the resource.
	   Possible values are: `SENSITIVITY_LOW`, `SENSITIVITY_MODERATE`, `SENSITIVITY_HIGH`.
	*/
	Score string `json:"score,omitempty" yaml:"score,omitempty"`
}
