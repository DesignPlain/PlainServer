package types

type Emr_InstanceGroupEbsConfig struct {
	// The number of I/O operations per second (IOPS) that the volume supports.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// The volume size, in gibibytes (GiB). This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The volume type. Valid options are 'gp2', 'io1' and 'standard'.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The number of EBS Volumes to attach per instance.
	VolumesPerInstance int `json:"volumesPerInstance,omitempty" yaml:"volumesPerInstance,omitempty"`
}
