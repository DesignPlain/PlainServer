package types

type S3_BucketReplicationConfigRuleSourceSelectionCriteria struct {
	// Configuration block that you can specify for selections for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when `filter` is specified), you can specify this element and set the status to `Enabled` to replicate modifications on replicas.
	ReplicaModifications S3_BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications `json:"replicaModifications,omitempty" yaml:"replicaModifications,omitempty"`

	// Configuration block for filter information for the selection of Amazon S3 objects encrypted with AWS KMS. If specified, `replica_kms_key_id` in `destination` `encryption_configuration` must be specified as well.
	SseKmsEncryptedObjects S3_BucketReplicationConfigRuleSourceSelectionCriteriaSseKmsEncryptedObjects `json:"sseKmsEncryptedObjects,omitempty" yaml:"sseKmsEncryptedObjects,omitempty"`
}
