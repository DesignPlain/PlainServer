package fsx

import types "DesignSphere_Server/src/resource/aws/types"

type OntapVolume struct {
	// Specifies the location in the storage virtual machine's namespace where the volume is mounted. The junction_path must have a leading forward slash, such as `/vol3`
	JunctionPath string `json:"junctionPath,omitempty" yaml:"junctionPath,omitempty"`

	// Specifies the size of the volume, in megabytes (MB), that you are creating.
	SizeInMegabytes int `json:"sizeInMegabytes,omitempty" yaml:"sizeInMegabytes,omitempty"`

	// The type of volume, currently the only valid value is `ONTAP`.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Setting this to `true` allows a SnapLock administrator to delete an FSx for ONTAP SnapLock Enterprise volume with unexpired write once, read many (WORM) files. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	BypassSnaplockEnterpriseRetention bool `json:"bypassSnaplockEnterpriseRetention,omitempty" yaml:"bypassSnaplockEnterpriseRetention,omitempty"`

	// Specifies the volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`.
	SecurityStyle string `json:"securityStyle,omitempty" yaml:"securityStyle,omitempty"`

	// When enabled, will skip the default final backup taken when the volume is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup bool `json:"skipFinalBackup,omitempty" yaml:"skipFinalBackup,omitempty"`

	// The SnapLock configuration for an FSx for ONTAP volume. See SnapLock Configuration below.
	SnaplockConfiguration types.Fsx_OntapVolumeSnaplockConfiguration `json:"snaplockConfiguration,omitempty" yaml:"snaplockConfiguration,omitempty"`

	// Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume.
	StorageEfficiencyEnabled bool `json:"storageEfficiencyEnabled,omitempty" yaml:"storageEfficiencyEnabled,omitempty"`

	// Specifies the storage virtual machine in which to create the volume.
	StorageVirtualMachineId string `json:"storageVirtualMachineId,omitempty" yaml:"storageVirtualMachineId,omitempty"`

	// A map of tags to assign to the volume. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The data tiering policy for an FSx for ONTAP volume. See Tiering Policy below.
	TieringPolicy types.Fsx_OntapVolumeTieringPolicy `json:"tieringPolicy,omitempty" yaml:"tieringPolicy,omitempty"`

	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A boolean flag indicating whether tags for the volume should be copied to backups. This value defaults to `false`.
	CopyTagsToBackups bool `json:"copyTagsToBackups,omitempty" yaml:"copyTagsToBackups,omitempty"`

	// Specifies the snapshot policy for the volume. See [snapshot policies](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snapshots-ontap.html#snapshot-policies) in the Amazon FSx ONTAP User Guide
	SnapshotPolicy string `json:"snapshotPolicy,omitempty" yaml:"snapshotPolicy,omitempty"`

	// Specifies the type of volume, valid values are `RW`, `DP`. Default value is `RW`. These can be set by the ONTAP CLI or API. This setting is used as part of migration and replication [Migrating to Amazon FSx for NetApp ONTAP](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/migrating-fsx-ontap.html)
	OntapVolumeType string `json:"ontapVolumeType,omitempty" yaml:"ontapVolumeType,omitempty"`
}
