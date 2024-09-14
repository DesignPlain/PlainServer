package types

type Compute_ResourcePolicySnapshotSchedulePolicyRetentionPolicy struct {
	/*
	   Specifies the behavior to apply to scheduled snapshots when
	   the source disk is deleted.
	   Default value is `KEEP_AUTO_SNAPSHOTS`.
	   Possible values are: `KEEP_AUTO_SNAPSHOTS`, `APPLY_RETENTION_POLICY`.
	*/
	OnSourceDiskDelete string `json:"onSourceDiskDelete,omitempty" yaml:"onSourceDiskDelete,omitempty"`

	// Maximum age of the snapshot that is allowed to be kept.
	MaxRetentionDays int `json:"maxRetentionDays,omitempty" yaml:"maxRetentionDays,omitempty"`
}
