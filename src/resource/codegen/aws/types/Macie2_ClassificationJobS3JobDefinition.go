package types

type Macie2_ClassificationJobS3JobDefinition struct {
	// The property- and tag-based conditions that determine which S3 buckets to include or exclude from the analysis. Conflicts with `bucket_definitions`. (documented below)
	BucketCriteria Macie2_ClassificationJobS3JobDefinitionBucketCriteria `json:"bucketCriteria,omitempty" yaml:"bucketCriteria,omitempty"`

	// An array of objects, one for each AWS account that owns buckets to analyze. Each object specifies the account ID for an account and one or more buckets to analyze for the account. Conflicts with `bucket_criteria`. (documented below)
	BucketDefinitions []Macie2_ClassificationJobS3JobDefinitionBucketDefinition `json:"bucketDefinitions,omitempty" yaml:"bucketDefinitions,omitempty"`

	// The property- and tag-based conditions that determine which objects to include or exclude from the analysis. (documented below)
	Scoping Macie2_ClassificationJobS3JobDefinitionScoping `json:"scoping,omitempty" yaml:"scoping,omitempty"`
}
