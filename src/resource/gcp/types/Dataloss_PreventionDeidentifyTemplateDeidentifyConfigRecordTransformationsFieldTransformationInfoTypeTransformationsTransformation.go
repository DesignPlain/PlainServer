package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformation struct {
	/*
	   InfoTypes to apply the transformation to. Leaving this empty will apply the transformation to apply to
	   all findings that correspond to infoTypes that were requested in InspectConfig.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformationInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`

	/*
	   Apply the transformation to the entire field.
	   The `primitive_transformation` block must only contain one argument, corresponding to the type of transformation.
	   Structure is documented below.
	*/
	PrimitiveTransformation Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformationPrimitiveTransformation `json:"primitiveTransformation,omitempty" yaml:"primitiveTransformation,omitempty"`
}
