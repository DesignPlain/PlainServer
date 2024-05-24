package types

type Sagemaker_FeatureGroupOfflineStoreConfigS3StorageConfig struct {
	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The S3 path where offline records are written.
	ResolvedOutputS3Uri string `json:"resolvedOutputS3Uri,omitempty" yaml:"resolvedOutputS3Uri,omitempty"`

	// The S3 URI, or location in Amazon S3, of OfflineStore.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`
}
