package types

type Ec2_LaunchConfigurationRootBlockDevice struct {
	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`
}
