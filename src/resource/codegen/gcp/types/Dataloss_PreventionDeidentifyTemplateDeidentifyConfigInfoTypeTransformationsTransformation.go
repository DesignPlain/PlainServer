package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformation struct {
	/*
	   Apply the transformation to the entire field.
	   The `primitive_transformation` block must only contain one argument, corresponding to the type of transformation.
	   Structure is documented below.
	*/
	PrimitiveTransformation Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformation `json:"primitiveTransformation,omitempty" yaml:"primitiveTransformation,omitempty"`

	/*
	   InfoTypes to apply the transformation to. Leaving this empty will apply the transformation to apply to
	   all findings that correspond to infoTypes that were requested in InspectConfig.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`
}
