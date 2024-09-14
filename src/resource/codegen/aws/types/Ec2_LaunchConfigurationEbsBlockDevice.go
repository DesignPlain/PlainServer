package types

type Ec2_LaunchConfigurationEbsBlockDevice struct {
	//
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	//
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`
}
