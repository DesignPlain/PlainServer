package types

type Rds_ClusterRestoreToPointInTime struct {
	// Date and time in UTC format to restore the database cluster to. Conflicts with `use_latest_restorable_time`.
	RestoreToTime string `json:"restoreToTime,omitempty" yaml:"restoreToTime,omitempty"`

	/*
	   Type of restore to be performed.
	   Valid options are `full-copy` (default) and `copy-on-write`.
	*/
	RestoreType string `json:"restoreType,omitempty" yaml:"restoreType,omitempty"`

	// Identifier of the source database cluster from which to restore. When restoring from a cluster in another AWS account, the identifier is the ARN of that cluster.
	SourceClusterIdentifier string `json:"sourceClusterIdentifier,omitempty" yaml:"sourceClusterIdentifier,omitempty"`

	// Cluster resource ID of the source database cluster from which to restore. To be used for restoring a deleted cluster in the same account which still has a retained automatic backup available.
	SourceClusterResourceId string `json:"sourceClusterResourceId,omitempty" yaml:"sourceClusterResourceId,omitempty"`

	// Set to true to restore the database cluster to the latest restorable backup time. Defaults to false. Conflicts with `restore_to_time`.
	UseLatestRestorableTime bool `json:"useLatestRestorableTime,omitempty" yaml:"useLatestRestorableTime,omitempty"`
}
