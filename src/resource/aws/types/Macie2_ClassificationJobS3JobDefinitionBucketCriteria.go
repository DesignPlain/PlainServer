package types

type Macie2_ClassificationJobS3JobDefinitionBucketCriteria struct {
	// The property- or tag-based conditions that determine which S3 buckets to exclude from the analysis. (documented below)
	Excludes Macie2_ClassificationJobS3JobDefinitionBucketCriteriaExcludes `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	// The property- or tag-based conditions that determine which S3 buckets to include in the analysis. (documented below)
	Includes Macie2_ClassificationJobS3JobDefinitionBucketCriteriaIncludes `json:"includes,omitempty" yaml:"includes,omitempty"`
}
