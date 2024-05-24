package types

type S3_BucketReplicationConfigRuleDestination struct {
	// Configuration block that specifies the overrides to use for object owners on replication. See below. Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object. Must be used in conjunction with `account` owner override configuration.
	AccessControlTranslation S3_BucketReplicationConfigRuleDestinationAccessControlTranslation `json:"accessControlTranslation,omitempty" yaml:"accessControlTranslation,omitempty"`

	// Account ID to specify the replica ownership. Must be used in conjunction with `access_control_translation` override configuration.
	Account string `json:"account,omitempty" yaml:"account,omitempty"`

	// ARN of the bucket where you want Amazon S3 to store the results.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Configuration block that provides information about encryption. See below. If `source_selection_criteria` is specified, you must specify this element.
	EncryptionConfiguration S3_BucketReplicationConfigRuleDestinationEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// Configuration block that specifies replication metrics-related settings enabling replication metrics and events. See below.
	Metrics S3_BucketReplicationConfigRuleDestinationMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	// Configuration block that specifies S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. See below. Replication Time Control must be used in conjunction with `metrics`.
	ReplicationTime S3_BucketReplicationConfigRuleDestinationReplicationTime `json:"replicationTime,omitempty" yaml:"replicationTime,omitempty"`

	// The [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Destination.html#AmazonS3-Type-Destination-StorageClass) used to store the object. By default, Amazon S3 uses the storage class of the source object to create the object replica.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}
