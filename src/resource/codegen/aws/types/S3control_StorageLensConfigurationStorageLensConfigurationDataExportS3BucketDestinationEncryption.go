package types

type S3control_StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryption struct {
	// SSE-KMS encryption. See SSE KMS below for more details.
	SseKms S3control_StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKms `json:"sseKms,omitempty" yaml:"sseKms,omitempty"`

	// SSE-S3 encryption. An empty configuration block `{}` should be used.
	SseS3s []S3control_StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3 `json:"sseS3s,omitempty" yaml:"sseS3s,omitempty"`
}
