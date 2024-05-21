package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformations struct {
	/*
	   Transformation for each infoType. Cannot specify more than one for a given infoType.
	   Structure is documented below.
	*/
	Transformations []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformation `json:"transformations,omitempty" yaml:"transformations,omitempty"`
}
