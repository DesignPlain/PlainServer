package types

type S3_BucketReplicationConfigurationRuleDestination struct {
	// The Account ID to use for overriding the object owner on replication. Must be used in conjunction with `access_control_translation` override configuration.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Enables replication metrics (required for S3 RTC) (documented below).
	Metrics S3_BucketReplicationConfigurationRuleDestinationMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	/*
	   Destination KMS encryption key ARN for SSE-KMS replication. Must be used in conjunction with
	   `sse_kms_encrypted_objects` source selection criteria.
	*/
	ReplicaKmsKeyId string `json:"replicaKmsKeyId,omitempty" yaml:"replicaKmsKeyId,omitempty"`

	// Enables S3 Replication Time Control (S3 RTC) (documented below).
	ReplicationTime S3_BucketReplicationConfigurationRuleDestinationReplicationTime `json:"replicationTime,omitempty" yaml:"replicationTime,omitempty"`

	// The [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Destination.html#AmazonS3-Type-Destination-StorageClass) used to store the object. By default, Amazon S3 uses the storage class of the source object to create the object replica.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// Specifies the overrides to use for object owners on replication. Must be used in conjunction with `account_id` owner override configuration.
	AccessControlTranslation S3_BucketReplicationConfigurationRuleDestinationAccessControlTranslation `json:"accessControlTranslation,omitempty" yaml:"accessControlTranslation,omitempty"`
}
