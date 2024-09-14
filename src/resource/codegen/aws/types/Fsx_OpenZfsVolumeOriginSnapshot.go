package types

type Fsx_OpenZfsVolumeOriginSnapshot struct {
	// Specifies the strategy used when copying data from the snapshot to the new volume. Valid values are `CLONE`, `FULL_COPY`, `INCREMENTAL_COPY`.
	CopyStrategy string `json:"copyStrategy,omitempty" yaml:"copyStrategy,omitempty"`

	// The Amazon Resource Name (ARN) of the origin snapshot.
	SnapshotArn string `json:"snapshotArn,omitempty" yaml:"snapshotArn,omitempty"`
}
