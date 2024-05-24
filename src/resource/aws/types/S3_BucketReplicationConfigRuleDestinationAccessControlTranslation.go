package types

type S3_BucketReplicationConfigRuleDestinationAccessControlTranslation struct {
	// Specifies the replica ownership. For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) in the Amazon S3 API Reference. Valid values: `Destination`.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`
}
