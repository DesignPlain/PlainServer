package types

type S3_BucketV2ReplicationConfigurationRuleDestinationAccessControlTranslation struct {
	// Specifies the replica ownership. For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html) in the Amazon S3 API Reference. The only valid value is `Destination`.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`
}
