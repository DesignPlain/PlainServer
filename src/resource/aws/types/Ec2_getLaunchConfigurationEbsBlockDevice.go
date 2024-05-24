package types

type Ec2_getLaunchConfigurationEbsBlockDevice struct {
	// Whether the EBS Volume will be deleted on instance termination.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Whether the device in the block device mapping of the AMI is suppressed.
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	// Throughput of the volume.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Size of the volume.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// Type of the volume.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Name of the device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Whether the volume is Encrypted.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Provisioned IOPs of the volume.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Snapshot ID of the mount.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
}
