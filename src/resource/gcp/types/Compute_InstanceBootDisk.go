package types

type Compute_InstanceBootDisk struct {
	/*
	   The name or self_link of the existing disk (such as those managed by
	   `gcp.compute.Disk`) or disk image. To create an instance from a snapshot, first create a
	   `gcp.compute.Disk` from a snapshot and reference it here.
	*/
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   Whether the disk will be auto-deleted when the instance
	   is deleted. Defaults to true.
	*/
	AutoDelete bool `json:"autoDelete,omitempty" yaml:"autoDelete,omitempty"`

	/*
	   Name with which attached disk will be accessible.
	   On the instance, this device will be `/dev/disk/by-id/google-{{device_name}}`.
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	/*
	   A 256-bit [customer-supplied encryption key]
	   (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
	   encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
	   to encrypt this disk. Only one of `kms_key_self_link` and `disk_encryption_key_raw`
	   may be set.
	*/
	DiskEncryptionKeyRaw string `json:"diskEncryptionKeyRaw,omitempty" yaml:"diskEncryptionKeyRaw,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 string `json:"diskEncryptionKeySha256,omitempty" yaml:"diskEncryptionKeySha256,omitempty"`

	/*
	   Parameters for a new disk that will be created
	   alongside the new instance. Either `initialize_params` or `source` must be set.
	   Structure is documented below.
	*/
	InitializeParams Compute_InstanceBootDiskInitializeParams `json:"initializeParams,omitempty" yaml:"initializeParams,omitempty"`

	/*
	   The self_link of the encryption key that is
	   stored in Google Cloud KMS to encrypt this disk. Only one of `kms_key_self_link`
	   and `disk_encryption_key_raw` may be set.
	*/
	KmsKeySelfLink string `json:"kmsKeySelfLink,omitempty" yaml:"kmsKeySelfLink,omitempty"`

	/*
	   The mode in which to attach this disk, either `READ_WRITE`
	   or `READ_ONLY`. If not specified, the default is to attach the disk in `READ_WRITE` mode.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
