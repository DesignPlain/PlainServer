package types



type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationFixedSizeBucketingConfig struct {
	/*
	   Size of each bucket (except for minimum and maximum buckets).
	   So if lower_bound = 10, upper_bound = 89, and bucketSize = 10, then the following buckets would be used: -10, 10-20, 20-30, 30-40, 40-50, 50-60, 60-70, 70-80, 80-89, 89+.
	   Precision up to 2 decimals works.
	*/
	BucketSize float64 `json:"bucketSize,omitempty" yaml:"bucketSize,omitempty"`

	/*
	   Lower bound value of buckets.
	   All values less than lower_bound are grouped together into a single bucket; for example if lower_bound = 10, then all values less than 10 are replaced with the value "-10".
	   The `lower_bound` block must only contain one argument. See the `fixed_size_bucketing_config` block description for more information about choosing a data type.
	   Structure is documented below.
	*/
	LowerBound Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationFixedSizeBucketingConfigLowerBound `json:"lowerBound,omitempty" yaml:"lowerBound,omitempty"`

	/*
	   Upper bound value of buckets.
	   All values greater than upper_bound are grouped together into a single bucket; for example if upper_bound = 89, then all values greater than 89 are replaced with the value "89+".
	   The `upper_bound` block must only contain one argument. See the `fixed_size_bucketing_config` block description for more information about choosing a data type.
	   Structure is documented below.
	*/
	UpperBound Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationFixedSizeBucketingConfigUpperBound `json:"upperBound,omitempty" yaml:"upperBound,omitempty"`
}
