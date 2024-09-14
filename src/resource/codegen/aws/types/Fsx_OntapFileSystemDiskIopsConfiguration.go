package types

type Fsx_OntapFileSystemDiskIopsConfiguration struct {
	// Specifies whether the number of IOPS for the file system is using the system. Valid values are `AUTOMATIC` and `USER_PROVISIONED`. Default value is `AUTOMATIC`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// The total number of SSD IOPS provisioned for the file system.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`
}
