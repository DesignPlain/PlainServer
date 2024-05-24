package types

type S3_BucketReplicationConfigRuleDestinationEncryptionConfiguration struct {
	// ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	ReplicaKmsKeyId string `json:"replicaKmsKeyId,omitempty" yaml:"replicaKmsKeyId,omitempty"`
}
