package types

type Docdb_ClusterRestoreToPointInTime struct {
	// A boolean value that indicates whether the DB cluster is restored from the latest backup time. Defaults to `false`. Cannot be specified with `restore_to_time`.
	UseLatestRestorableTime bool `json:"useLatestRestorableTime,omitempty" yaml:"useLatestRestorableTime,omitempty"`

	// The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be before the latest restorable time for the DB instance. Cannot be specified with `use_latest_restorable_time`.
	RestoreToTime string `json:"restoreToTime,omitempty" yaml:"restoreToTime,omitempty"`

	// The type of restore to be performed. Valid values are `full-copy`, `copy-on-write`.
	RestoreType string `json:"restoreType,omitempty" yaml:"restoreType,omitempty"`

	// The identifier of the source DB cluster from which to restore. Must match the identifier of an existing DB cluster.
	SourceClusterIdentifier string `json:"sourceClusterIdentifier,omitempty" yaml:"sourceClusterIdentifier,omitempty"`
}
