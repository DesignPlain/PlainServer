package types

type Ec2_LaunchConfigurationEbsBlockDevice struct {
	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	//
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	//
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`
}
