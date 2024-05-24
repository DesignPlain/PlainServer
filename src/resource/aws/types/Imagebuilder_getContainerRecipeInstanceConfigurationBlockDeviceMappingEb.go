package types

type Imagebuilder_getContainerRecipeInstanceConfigurationBlockDeviceMappingEb struct {
	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Number of Input/Output (I/O) operations per second to provision for an `io1` or `io2` volume.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// KMS key used to encrypt the container image.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Identifier of the EC2 Volume Snapshot.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	// For GP3 volumes only. The throughput in MiB/s that the volume supports.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Size of the volume, in GiB.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// Type of the volume. For example, `gp2` or `io2`.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`
}
