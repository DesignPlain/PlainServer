package types

type Fsx_FileCacheLustreConfigurationMetadataConfiguration struct {
	// The storage capacity of the Lustre MDT (Metadata Target) storage volume in gibibytes (GiB). The only supported value is `2400` GiB.
	StorageCapacity int `json:"storageCapacity,omitempty" yaml:"storageCapacity,omitempty"`
}
