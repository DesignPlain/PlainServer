package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformation struct {
	/*
	   Only apply the transformation if the condition evaluates to true for the given RecordCondition. The conditions are allowed to reference fields that are not used in the actual transformation.
	   Example Use Cases:
	   - Apply a different bucket transformation to an age column if the zip code column for the same record is within a specific range.
	   - Redact a field if the date of birth field is greater than 85.
	   Structure is documented below.
	*/
	Condition Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	/*
	   Input field(s) to apply the transformation to. When you have columns that reference their position within a list, omit the index from the FieldId.
	   FieldId name matching ignores the index. For example, instead of "contact.nums[0].type", use "contact.nums.type".
	   Structure is documented below.
	*/
	Fields []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationField `json:"fields,omitempty" yaml:"fields,omitempty"`

	/*
	   Treat the contents of the field as free text, and selectively transform content that matches an InfoType.
	   Only one of `primitive_transformation` or `info_type_transformations` must be specified.
	   Structure is documented below.
	*/
	InfoTypeTransformations Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformations `json:"infoTypeTransformations,omitempty" yaml:"infoTypeTransformations,omitempty"`

	/*
	   Apply the transformation to the entire field.
	   The `primitive_transformation` block must only contain one argument, corresponding to the type of transformation.
	   Only one of `primitive_transformation` or `info_type_transformations` must be specified.
	   Structure is documented below.
	*/
	PrimitiveTransformation Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformation `json:"primitiveTransformation,omitempty" yaml:"primitiveTransformation,omitempty"`
}
