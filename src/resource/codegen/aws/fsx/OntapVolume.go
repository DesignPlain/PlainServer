package fsx

import types "libds/aws/types"

type OntapVolume struct {
	// The Aggregate configuration only applies to `FLEXGROUP` volumes. See [`aggregate_configuration` Block] for details.
	AggregateConfiguration types.Fsx_OntapVolumeAggregateConfiguration `json:"aggregateConfiguration,omitempty" yaml:"aggregateConfiguration,omitempty"`

	// Specifies the location in the storage virtual machine's namespace where the volume is mounted. The junction_path must have a leading forward slash, such as `/vol3`
	JunctionPath string `json:"junctionPath,omitempty" yaml:"junctionPath,omitempty"`

	// When enabled, will skip the default final backup taken when the volume is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup bool `json:"skipFinalBackup,omitempty" yaml:"skipFinalBackup,omitempty"`

	// Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume.
	StorageEfficiencyEnabled bool `json:"storageEfficiencyEnabled,omitempty" yaml:"storageEfficiencyEnabled,omitempty"`

	// Setting this to `true` allows a SnapLock administrator to delete an FSx for ONTAP SnapLock Enterprise volume with unexpired write once, read many (WORM) files. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	BypassSnaplockEnterpriseRetention bool `json:"bypassSnaplockEnterpriseRetention,omitempty" yaml:"bypassSnaplockEnterpriseRetention,omitempty"`

	// Specifies the snapshot policy for the volume. See [snapshot policies](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snapshots-ontap.html#snapshot-policies) in the Amazon FSx ONTAP User Guide
	SnapshotPolicy string `json:"snapshotPolicy,omitempty" yaml:"snapshotPolicy,omitempty"`

	// The data tiering policy for an FSx for ONTAP volume. See `tiering_policy` Block for details.
	TieringPolicy types.Fsx_OntapVolumeTieringPolicy `json:"tieringPolicy,omitempty" yaml:"tieringPolicy,omitempty"`

	// Specifies the styles of volume, valid values are `FLEXVOL`, `FLEXGROUP`. Default value is `FLEXVOL`. FLEXGROUPS have a larger minimum and maximum size. See Volume Styles for more details. [Volume Styles](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-styles.html)
	VolumeStyle string `json:"volumeStyle,omitempty" yaml:"volumeStyle,omitempty"`

	// Specifies the size of the volume, in megabytes (MB), that you are creating. Supported when creating volumes under 2 PB. Either size_in_bytes or size_in_megabytes must be specified. Minimum size for `FLEXGROUP` volumes are 100GiB per constituent.
	SizeInMegabytes int `json:"sizeInMegabytes,omitempty" yaml:"sizeInMegabytes,omitempty"`

	// The SnapLock configuration for an FSx for ONTAP volume. See `snaplock_configuration` Block for details.
	SnaplockConfiguration types.Fsx_OntapVolumeSnaplockConfiguration `json:"snaplockConfiguration,omitempty" yaml:"snaplockConfiguration,omitempty"`

	// A map of tags to assign to the volume. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of volume, currently the only valid value is `ONTAP`.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// A boolean flag indicating whether tags for the volume should be copied to backups. This value defaults to `false`.
	CopyTagsToBackups bool `json:"copyTagsToBackups,omitempty" yaml:"copyTagsToBackups,omitempty"`

	// A map of tags to apply to the volume's final backup.
	FinalBackupTags map[string]string `json:"finalBackupTags,omitempty" yaml:"finalBackupTags,omitempty"`

	// Specifies the type of volume, valid values are `RW`, `DP`. Default value is `RW`. These can be set by the ONTAP CLI or API. This setting is used as part of migration and replication [Migrating to Amazon FSx for NetApp ONTAP](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/migrating-fsx-ontap.html)
	OntapVolumeType string `json:"ontapVolumeType,omitempty" yaml:"ontapVolumeType,omitempty"`

	// Specifies the size of the volume, in megabytes (MB), that you are creating. Can be used for any size but required for volumes over 2 PB. Either size_in_bytes or size_in_megabytes must be specified. Minimum size for `FLEXGROUP` volumes are 100GiB per constituent.
	SizeInBytes string `json:"sizeInBytes,omitempty" yaml:"sizeInBytes,omitempty"`

	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`.
	SecurityStyle string `json:"securityStyle,omitempty" yaml:"securityStyle,omitempty"`

	/*
	   Specifies the storage virtual machine in which to create the volume.

	   The following arguments are optional:
	*/
	StorageVirtualMachineId string `json:"storageVirtualMachineId,omitempty" yaml:"storageVirtualMachineId,omitempty"`
}
