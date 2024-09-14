package types

type Alloydb_ClusterMigrationSource struct {
	// The host and port of the on-premises instance in host:port format
	HostPort string `json:"hostPort,omitempty" yaml:"hostPort,omitempty"`

	// Place holder for the external source identifier(e.g DMS job name) that created the cluster.
	ReferenceId string `json:"referenceId,omitempty" yaml:"referenceId,omitempty"`

	// Type of migration source.
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`
}
