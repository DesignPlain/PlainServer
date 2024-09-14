package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformations struct {
	/*
	   For determination of how redaction of images should occur.
	   Structure is documented below.
	*/
	Transforms []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform `json:"transforms,omitempty" yaml:"transforms,omitempty"`
}
