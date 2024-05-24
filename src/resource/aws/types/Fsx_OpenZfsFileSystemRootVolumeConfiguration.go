package types

type Fsx_OpenZfsFileSystemRootVolumeConfiguration struct {
	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
	UserAndGroupQuotas []Fsx_OpenZfsFileSystemRootVolumeConfigurationUserAndGroupQuota `json:"userAndGroupQuotas,omitempty" yaml:"userAndGroupQuotas,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots bool `json:"copyTagsToSnapshots,omitempty" yaml:"copyTagsToSnapshots,omitempty"`

	// Method used to compress the data on the volume. Valid values are `LZ4`, `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType string `json:"dataCompressionType,omitempty" yaml:"dataCompressionType,omitempty"`

	// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
	NfsExports Fsx_OpenZfsFileSystemRootVolumeConfigurationNfsExports `json:"nfsExports,omitempty" yaml:"nfsExports,omitempty"`

	// specifies whether the volume is read-only. Default is false.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// Specifies the record size of an OpenZFS root volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib int `json:"recordSizeKib,omitempty" yaml:"recordSizeKib,omitempty"`
}
