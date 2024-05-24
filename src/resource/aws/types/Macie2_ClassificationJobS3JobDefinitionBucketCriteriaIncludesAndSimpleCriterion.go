package types

type Macie2_ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndSimpleCriterion struct {
	// The object property to use in the condition. Valid combination of values are available in the [AWS Documentation](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html#jobs-model-simplecriterionkeyforjob)
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// An array that lists the values to use in the condition. Valid combination of values are available in the [AWS Documentation](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html#jobs-model-simplecriterionforjob)
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The operator to use in a condition. Valid combination of values are available in the [AWS Documentation](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html#jobs-model-jobcomparator)
	Comparator string `json:"comparator,omitempty" yaml:"comparator,omitempty"`
}
