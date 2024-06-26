package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationTimePartConfig struct {
	/*
	   The part of the time to keep.
	   Possible values are: `YEAR`, `MONTH`, `DAY_OF_MONTH`, `DAY_OF_WEEK`, `WEEK_OF_YEAR`, `HOUR_OF_DAY`.
	*/
	PartToExtract string `json:"partToExtract,omitempty" yaml:"partToExtract,omitempty"`
}
