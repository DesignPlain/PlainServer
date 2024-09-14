package types

type Macie2_ClassificationJobS3JobDefinitionBucketCriteriaIncludes struct {
	// An array of conditions, one for each condition that determines which objects to include or exclude from the job. (documented below)
	Ands []Macie2_ClassificationJobS3JobDefinitionBucketCriteriaIncludesAnd `json:"ands,omitempty" yaml:"ands,omitempty"`
}
