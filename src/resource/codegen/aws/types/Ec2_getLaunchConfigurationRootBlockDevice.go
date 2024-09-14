package types

type Ec2_getLaunchConfigurationRootBlockDevice struct {
	// Whether the EBS Volume will be deleted on instance termination.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Whether the volume is Encrypted.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Provisioned IOPs of the volume.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Throughput of the volume.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Size of the volume.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// Type of the volume.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`
}
