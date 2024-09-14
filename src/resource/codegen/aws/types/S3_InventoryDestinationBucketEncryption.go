package types

type S3_InventoryDestinationBucketEncryption struct {
	// Specifies to use server-side encryption with Amazon S3-managed keys (SSE-S3) to encrypt the inventory file.
	SseS3 S3_InventoryDestinationBucketEncryptionSseS3 `json:"sseS3,omitempty" yaml:"sseS3,omitempty"`

	// Specifies to use server-side encryption with AWS KMS-managed keys to encrypt the inventory file (documented below).
	SseKms S3_InventoryDestinationBucketEncryptionSseKms `json:"sseKms,omitempty" yaml:"sseKms,omitempty"`
}
