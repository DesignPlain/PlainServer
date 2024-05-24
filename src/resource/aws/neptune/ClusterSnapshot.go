package neptune

type ClusterSnapshot struct {
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier string `json:"dbClusterIdentifier,omitempty" yaml:"dbClusterIdentifier,omitempty"`

	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier,omitempty" yaml:"dbClusterSnapshotIdentifier,omitempty"`
}
