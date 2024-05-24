package types

type S3_BucketV2ReplicationConfigurationRuleSourceSelectionCriteria struct {
	/*
	   Match SSE-KMS encrypted objects (documented below). If specified, `replica_kms_key_id`
	   in `destination` must be specified as well.
	*/
	SseKmsEncryptedObjects []S3_BucketV2ReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObject `json:"sseKmsEncryptedObjects,omitempty" yaml:"sseKmsEncryptedObjects,omitempty"`
}
