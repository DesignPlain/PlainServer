package fsx

import types "libds/aws/types"

type OpenZfsVolume struct {
	// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib int `json:"recordSizeKib,omitempty" yaml:"recordSizeKib,omitempty"`

	// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
	StorageCapacityQuotaGib int `json:"storageCapacityQuotaGib,omitempty" yaml:"storageCapacityQuotaGib,omitempty"`

	// A map of tags to assign to the file system. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See `user_and_group_quotas` Block Below.
	UserAndGroupQuotas []types.Fsx_OpenZfsVolumeUserAndGroupQuota `json:"userAndGroupQuotas,omitempty" yaml:"userAndGroupQuotas,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots bool `json:"copyTagsToSnapshots,omitempty" yaml:"copyTagsToSnapshots,omitempty"`

	// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType string `json:"dataCompressionType,omitempty" yaml:"dataCompressionType,omitempty"`

	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
	StorageCapacityReservationGib int `json:"storageCapacityReservationGib,omitempty" yaml:"storageCapacityReservationGib,omitempty"`

	//
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// NFS export configuration for the root volume. Exactly 1 item. See `nfs_exports` Block Below for details.
	NfsExports types.Fsx_OpenZfsVolumeNfsExports `json:"nfsExports,omitempty" yaml:"nfsExports,omitempty"`

	// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `aws.fsx.OpenZfsFileSystem` resource with the `root_volume_id` or the `id` property of another `aws.fsx.OpenZfsVolume`.
	ParentVolumeId string `json:"parentVolumeId,omitempty" yaml:"parentVolumeId,omitempty"`

	// specifies whether the volume is read-only. Default is false.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// Whether to delete all child volumes and snapshots. Valid values: `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS`. This configuration must be applied separately before attempting to delete the resource to have the desired behavior..
	DeleteVolumeOptions string `json:"deleteVolumeOptions,omitempty" yaml:"deleteVolumeOptions,omitempty"`

	// Specifies the configuration to use when creating the OpenZFS volume. See `origin_snapshot` Block below for details.
	OriginSnapshot types.Fsx_OpenZfsVolumeOriginSnapshot `json:"originSnapshot,omitempty" yaml:"originSnapshot,omitempty"`
}
