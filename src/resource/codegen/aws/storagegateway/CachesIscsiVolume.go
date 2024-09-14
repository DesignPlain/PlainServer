package storagegateway

type CachesIscsiVolume struct {
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`

	// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
	KmsEncrypted bool `json:"kmsEncrypted,omitempty" yaml:"kmsEncrypted,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kms_encrypted` is set.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName string `json:"targetName,omitempty" yaml:"targetName,omitempty"`

	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volume_size_in_bytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn string `json:"sourceVolumeArn,omitempty" yaml:"sourceVolumeArn,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The size of the volume in bytes.
	VolumeSizeInBytes int `json:"volumeSizeInBytes,omitempty" yaml:"volumeSizeInBytes,omitempty"`
}
