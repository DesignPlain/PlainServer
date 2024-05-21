package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationBucketingConfigBucket struct {
	/*
	   Upper bound of the range, exclusive; type must match min.
	   The `max` block must only contain one argument. See the `bucketing_config` block description for more information about choosing a data type.
	   Structure is documented below.
	*/
	Max Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationBucketingConfigBucketMax `json:"max,omitempty" yaml:"max,omitempty"`

	/*
	   Lower bound of the range, inclusive. Type should be the same as max if used.
	   The `min` block must only contain one argument. See the `bucketing_config` block description for more information about choosing a data type.
	   Structure is documented below.
	*/
	Min Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationBucketingConfigBucketMin `json:"min,omitempty" yaml:"min,omitempty"`

	/*
	   Replacement value for this bucket.
	   The `replacement_value` block must only contain one argument.
	   Structure is documented below.
	*/
	ReplacementValue Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationPrimitiveTransformationBucketingConfigBucketReplacementValue `json:"replacementValue,omitempty" yaml:"replacementValue,omitempty"`
}
