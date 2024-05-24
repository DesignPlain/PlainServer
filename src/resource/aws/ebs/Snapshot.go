package ebs

type Snapshot struct {
	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier string `json:"storageTier,omitempty" yaml:"storageTier,omitempty"`

	// A map of tags to assign to the snapshot. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays int `json:"temporaryRestoreDays,omitempty" yaml:"temporaryRestoreDays,omitempty"`

	// The Volume ID of which to make a snapshot.
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`

	// A description of what the snapshot is.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost on which to create a local snapshot.
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore bool `json:"permanentRestore,omitempty" yaml:"permanentRestore,omitempty"`
}
