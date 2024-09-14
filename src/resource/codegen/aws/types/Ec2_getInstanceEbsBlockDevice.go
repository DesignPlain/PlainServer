package types

type Ec2_getInstanceEbsBlockDevice struct {
	// ID of the snapshot.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`

	// Size of the volume, in GiB.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// If the root block device will be deleted on termination.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Physical name of the device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// `0` If the volume is not a provisioned IOPS image, otherwise the supported IOPS count.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// If the EBS volume is encrypted.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Map of tags assigned to the Instance.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Throughput of the volume, in MiB/s.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Type of the volume.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`
}
