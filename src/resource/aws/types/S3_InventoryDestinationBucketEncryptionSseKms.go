package types

type S3_InventoryDestinationBucketEncryptionSseKms struct {
	// ARN of the KMS customer master key (CMK) used to encrypt the inventory file.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`
}
