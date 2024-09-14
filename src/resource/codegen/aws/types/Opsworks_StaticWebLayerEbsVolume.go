package types

type Opsworks_StaticWebLayerEbsVolume struct {
	// The path to mount the EBS volume on the layer's instances.
	MountPoint string `json:"mountPoint,omitempty" yaml:"mountPoint,omitempty"`

	// The number of disks to use for the EBS volume.
	NumberOfDisks int `json:"numberOfDisks,omitempty" yaml:"numberOfDisks,omitempty"`

	// The RAID level to use for the volume.
	RaidLevel string `json:"raidLevel,omitempty" yaml:"raidLevel,omitempty"`

	// The size of the volume in gigabytes.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The type of volume to create. This may be `standard` (the default), `io1` or `gp2`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`
}
