package types

type Finspace_KxClusterDatabase struct {
	// The name of the dataview to be used for caching historical data on disk. You cannot update to a different dataview name once a cluster is created. Use `lifecycle` `ignore_changes` for database to prevent any undesirable behaviors.
	DataviewName string `json:"dataviewName,omitempty" yaml:"dataviewName,omitempty"`

	// Configuration details for the disk cache to increase performance reading from a KX database mounted to the cluster. See cache_configurations.
	CacheConfigurations []Finspace_KxClusterDatabaseCacheConfiguration `json:"cacheConfigurations,omitempty" yaml:"cacheConfigurations,omitempty"`

	// A unique identifier of the changeset that is associated with the cluster.
	ChangesetId string `json:"changesetId,omitempty" yaml:"changesetId,omitempty"`

	// Name of the KX database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
