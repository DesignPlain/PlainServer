package types

type Ec2_SpotFleetRequestLaunchSpecificationEbsBlockDevice struct {
	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	//
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	//
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
