package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformationPrimitiveTransformationBucketingConfig struct {
	/*
	   Set of buckets. Ranges must be non-overlapping.
	   Bucket is represented as a range, along with replacement values.
	   Structure is documented below.
	*/
	Buckets []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationInfoTypeTransformationsTransformationPrimitiveTransformationBucketingConfigBucket `json:"buckets,omitempty" yaml:"buckets,omitempty"`
}
