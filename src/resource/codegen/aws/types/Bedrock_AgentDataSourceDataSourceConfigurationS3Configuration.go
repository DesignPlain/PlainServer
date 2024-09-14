package types

type Bedrock_AgentDataSourceDataSourceConfigurationS3Configuration struct {
	// List of S3 prefixes that define the object containing the data sources. For more information, see [Organizing objects using prefixes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-prefixes.html).
	InclusionPrefixes []string `json:"inclusionPrefixes,omitempty" yaml:"inclusionPrefixes,omitempty"`

	// ARN of the bucket that contains the data source.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// Bucket account owner ID for the S3 bucket.
	BucketOwnerAccountId string `json:"bucketOwnerAccountId,omitempty" yaml:"bucketOwnerAccountId,omitempty"`
}
