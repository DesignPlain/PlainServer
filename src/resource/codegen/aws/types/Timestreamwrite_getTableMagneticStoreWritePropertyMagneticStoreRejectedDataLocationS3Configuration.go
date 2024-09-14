package types

type Timestreamwrite_getTableMagneticStoreWritePropertyMagneticStoreRejectedDataLocationS3Configuration struct {
	// AWS KMS key ID for S3 location with AWS maanged key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Object key preview for S3 location.
	ObjectKeyPrefix string `json:"objectKeyPrefix,omitempty" yaml:"objectKeyPrefix,omitempty"`

	// Name of S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	//
	EncryptionOption string `json:"encryptionOption,omitempty" yaml:"encryptionOption,omitempty"`
}
