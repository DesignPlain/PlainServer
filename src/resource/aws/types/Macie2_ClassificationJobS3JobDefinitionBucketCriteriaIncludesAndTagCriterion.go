package types

type Macie2_ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndTagCriterion struct {
	// The  tag key and value pairs to use in the condition. One or more blocks are allowed. (documented below)
	TagValues []Macie2_ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndTagCriterionTagValue `json:"tagValues,omitempty" yaml:"tagValues,omitempty"`

	// The operator to use in the condition. Valid combination and values are available in the [AWS Documentation](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html#jobs-model-jobcomparator)
	Comparator string `json:"comparator,omitempty" yaml:"comparator,omitempty"`
}
