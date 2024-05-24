package types

type Opsworks_InstanceRootBlockDevice struct {
	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`
}
