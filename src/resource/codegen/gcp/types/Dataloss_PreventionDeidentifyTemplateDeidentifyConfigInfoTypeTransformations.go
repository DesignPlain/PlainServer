package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations struct {
	/*
	   Transformation for each infoType. Cannot specify more than one for a given infoType.
	   Structure is documented below.
	*/
	Transformations []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformation `json:"transformations,omitempty" yaml:"transformations,omitempty"`
}
