package redshiftserverless

type Snapshot struct {
	// The namespace to create a snapshot for.
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`

	// How long to retain the created snapshot. Default value is `-1`.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	// The name of the snapshot.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`
}
