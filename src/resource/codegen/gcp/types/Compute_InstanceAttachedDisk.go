package types

type Compute_InstanceAttachedDisk struct {
	/*
	   Name with which the attached disk will be accessible
	   under `/dev/disk/by-id/google--`
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	/*
	   A 256-bit [customer-supplied encryption key]
	   (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
	   encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
	   to encrypt this disk. Only one of `kms_key_self_link` and `disk_encryption_key_raw` may be set.
	*/
	DiskEncryptionKeyRaw string `json:"diskEncryptionKeyRaw,omitempty" yaml:"diskEncryptionKeyRaw,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 string `json:"diskEncryptionKeySha256,omitempty" yaml:"diskEncryptionKeySha256,omitempty"`

	/*
	   The self_link of the encryption key that is
	   stored in Google Cloud KMS to encrypt this disk. Only one of `kms_key_self_link`
	   and `disk_encryption_key_raw` may be set.
	*/
	KmsKeySelfLink string `json:"kmsKeySelfLink,omitempty" yaml:"kmsKeySelfLink,omitempty"`

	/*
	   Either "READ_ONLY" or "READ_WRITE", defaults to "READ_WRITE"
	   If you have a persistent disk with data that you want to share
	   between multiple instances, detach it from any read-write instances and
	   attach it to one or more instances in read-only mode.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// The name or self_link of the disk to attach to this instance.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
