package types

type Macie2_ClassificationJobS3JobDefinitionBucketCriteriaExcludesAnd struct {
	// A property-based condition that defines a property, operator, and one or more values for including or excluding an S3 buckets from the job. (documented below)
	SimpleCriterion Macie2_ClassificationJobS3JobDefinitionBucketCriteriaExcludesAndSimpleCriterion `json:"simpleCriterion,omitempty" yaml:"simpleCriterion,omitempty"`

	// A tag-based condition that defines the operator and tag keys or tag key and value pairs for including or excluding an S3 buckets from the job. (documented below)
	TagCriterion Macie2_ClassificationJobS3JobDefinitionBucketCriteriaExcludesAndTagCriterion `json:"tagCriterion,omitempty" yaml:"tagCriterion,omitempty"`
}
