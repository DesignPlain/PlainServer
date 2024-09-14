package types

type Connect_InstanceStorageConfigStorageConfigS3Config struct {
	// The S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// The encryption configuration. Documented below.
	EncryptionConfig Connect_InstanceStorageConfigStorageConfigS3ConfigEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
}
