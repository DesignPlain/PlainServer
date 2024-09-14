package types

type Timestreamwrite_TableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3Configuration struct {
	// Encryption option for the customer s3 location. Options are S3 server side encryption with an S3-managed key or KMS managed key. Valid values are `SSE_KMS` and `SSE_S3`.
	EncryptionOption string `json:"encryptionOption,omitempty" yaml:"encryptionOption,omitempty"`

	// KMS key arn for the customer s3 location when encrypting with a KMS managed key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Object key prefix for the customer S3 location.
	ObjectKeyPrefix string `json:"objectKeyPrefix,omitempty" yaml:"objectKeyPrefix,omitempty"`

	// Bucket name of the customer S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
