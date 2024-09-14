package types

type Ecs_ServiceVolumeConfigurationManagedEbsVolume struct {
	// Whether the volume should be encrypted. Default value is `true`.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Linux filesystem type for the volume. For volumes created from a snapshot, same filesystem type must be specified that the volume was using when the snapshot was created. Valid values are `ext3`, `ext4`, `xfs`. Default value is `xfs`.
	FileSystemType string `json:"fileSystemType,omitempty" yaml:"fileSystemType,omitempty"`

	// Number of I/O operations per second (IOPS).
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Amazon Resource Name (ARN) identifier of the Amazon Web Services Key Management Service key to use for Amazon EBS encryption.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Size of the volume in GiB. You must specify either a `size_in_gb` or a `snapshot_id`. You can optionally specify a volume size greater than or equal to the snapshot size.
	SizeInGb int `json:"sizeInGb,omitempty" yaml:"sizeInGb,omitempty"`

	// Snapshot that Amazon ECS uses to create the volume. You must specify either a `size_in_gb` or a `snapshot_id`.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	// Throughput to provision for a volume, in MiB/s, with a maximum of 1,000 MiB/s.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Amazon ECS infrastructure IAM role that is used to manage your Amazon Web Services infrastructure. Recommended using the Amazon ECS-managed `AmazonECSInfrastructureRolePolicyForVolumes` IAM policy with this role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Volume type.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`
}
