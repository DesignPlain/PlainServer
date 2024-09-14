package ebs

type SnapshotCopy struct {
	// A description of what the snapshot is.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether the snapshot is encrypted.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// The ARN for the snapshot to be copied.
	SourceSnapshotId string `json:"sourceSnapshotId,omitempty" yaml:"sourceSnapshotId,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays int `json:"temporaryRestoreDays,omitempty" yaml:"temporaryRestoreDays,omitempty"`

	// The ARN for the KMS encryption key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore bool `json:"permanentRestore,omitempty" yaml:"permanentRestore,omitempty"`

	// The region of the source snapshot.
	SourceRegion string `json:"sourceRegion,omitempty" yaml:"sourceRegion,omitempty"`

	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier string `json:"storageTier,omitempty" yaml:"storageTier,omitempty"`

	// A map of tags for the snapshot.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
