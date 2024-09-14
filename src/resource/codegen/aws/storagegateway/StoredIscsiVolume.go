package storagegateway

type StoredIscsiVolume struct {
	// The unique identifier for the gateway local disk that is configured as a stored volume.
	DiskId string `json:"diskId,omitempty" yaml:"diskId,omitempty"`

	// Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
	PreserveExistingData bool `json:"preserveExistingData,omitempty" yaml:"preserveExistingData,omitempty"`

	// The snapshot ID of the snapshot to restore as the new stored volumeE.g., `snap-1122aabb`.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	// Key-value mapping of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName string `json:"targetName,omitempty" yaml:"targetName,omitempty"`

	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`

	// `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
	KmsEncrypted bool `json:"kmsEncrypted,omitempty" yaml:"kmsEncrypted,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is `true`.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`
}
