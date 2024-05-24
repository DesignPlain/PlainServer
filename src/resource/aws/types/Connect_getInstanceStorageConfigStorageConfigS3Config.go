package types

type Connect_getInstanceStorageConfigStorageConfigS3Config struct {
	// The S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// The encryption configuration. Documented below.
	EncryptionConfigs []Connect_getInstanceStorageConfigStorageConfigS3ConfigEncryptionConfig `json:"encryptionConfigs,omitempty" yaml:"encryptionConfigs,omitempty"`
}
