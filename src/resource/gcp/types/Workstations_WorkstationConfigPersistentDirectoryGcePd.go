package types

type Workstations_WorkstationConfigPersistentDirectoryGcePd struct {
	// The type of the persistent disk for the home directory. Defaults to `pd-standard`.
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	// Type of file system that the disk should be formatted with. The workstation image must support this file system type. Must be empty if `sourceSnapshot` is set. Defaults to `ext4`.
	FsType string `json:"fsType,omitempty" yaml:"fsType,omitempty"`

	/*
	   Whether the persistent disk should be deleted when the workstation is deleted. Valid values are `DELETE` and `RETAIN`. Defaults to `DELETE`.
	   Possible values are: `DELETE`, `RETAIN`.
	*/
	ReclaimPolicy string `json:"reclaimPolicy,omitempty" yaml:"reclaimPolicy,omitempty"`

	/*
	   The GB capacity of a persistent home directory for each workstation created with this configuration. Must be empty if `sourceSnapshot` is set.
	   Valid values are `10`, `50`, `100`, `200`, `500`, or `1000`. Defaults to `200`. If less than `200` GB, the `diskType` must be `pd-balanced` or `pd-ssd`.
	*/
	SizeGb int `json:"sizeGb,omitempty" yaml:"sizeGb,omitempty"`

	// Name of the snapshot to use as the source for the disk. This can be the snapshot's `self_link`, `id`, or a string in the format of `projects/{project}/global/snapshots/{snapshot}`. If set, `sizeGb` and `fsType` must be empty. Can only be updated if it has an existing value.
	SourceSnapshot string `json:"sourceSnapshot,omitempty" yaml:"sourceSnapshot,omitempty"`
}
