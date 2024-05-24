package types

type Emr_InstanceFleetInstanceTypeConfigEbsConfig struct {
	// The number of I/O operations per second (IOPS) that the volume supports
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// The volume size, in gibibytes (GiB).
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The volume type. Valid options are `gp2`, `io1`, `standard` and `st1`. See [EBS Volume Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumehtml).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The number of EBS volumes with this configuration to attach to each EC2 instance in the instance group (default is 1)
	VolumesPerInstance int `json:"volumesPerInstance,omitempty" yaml:"volumesPerInstance,omitempty"`
}
