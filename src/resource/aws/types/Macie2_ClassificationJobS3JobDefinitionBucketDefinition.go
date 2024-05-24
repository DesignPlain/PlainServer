package types

type Macie2_ClassificationJobS3JobDefinitionBucketDefinition struct {
	// The unique identifier for the AWS account that owns the buckets.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// An array that lists the names of the buckets.
	Buckets []string `json:"buckets,omitempty" yaml:"buckets,omitempty"`
}
