package types

type Ec2_getLaunchTemplateBlockDeviceMappingEb struct {
	//
	Encrypted string `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	//
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	//
	DeleteOnTermination string `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`
}
