package types

type Opsworks_InstanceEbsBlockDevice struct {
	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`
}
