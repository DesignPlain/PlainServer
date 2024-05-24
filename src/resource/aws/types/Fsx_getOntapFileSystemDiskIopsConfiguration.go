package types

type Fsx_getOntapFileSystemDiskIopsConfiguration struct {
	// The total number of SSD IOPS provisioned for the file system.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Specifies whether the file system is using the `AUTOMATIC` setting of SSD IOPS of 3 IOPS per GB of storage capacity, or if it using a `USER_PROVISIONED` value.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
