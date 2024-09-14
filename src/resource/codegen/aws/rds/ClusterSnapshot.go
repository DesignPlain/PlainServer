package rds

type ClusterSnapshot struct {
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier string `json:"dbClusterIdentifier,omitempty" yaml:"dbClusterIdentifier,omitempty"`

	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier,omitempty" yaml:"dbClusterSnapshotIdentifier,omitempty"`

	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts []string `json:"sharedAccounts,omitempty" yaml:"sharedAccounts,omitempty"`

	// A map of tags to assign to the DB cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
