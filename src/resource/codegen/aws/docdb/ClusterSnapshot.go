package docdb

type ClusterSnapshot struct {
	// The DocumentDB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier string `json:"dbClusterIdentifier,omitempty" yaml:"dbClusterIdentifier,omitempty"`

	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier,omitempty" yaml:"dbClusterSnapshotIdentifier,omitempty"`
}
