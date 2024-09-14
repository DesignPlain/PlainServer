package docdb

type GlobalCluster struct {
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `docdb`. Defaults to `docdb`. Conflicts with `source_db_cluster_identifier`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	/*
	   Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
	   - --NOTE:-- Upgrading major versions is not supported.
	*/
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// The global cluster identifier.
	GlobalClusterIdentifier string `json:"globalClusterIdentifier,omitempty" yaml:"globalClusterIdentifier,omitempty"`

	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier string `json:"sourceDbClusterIdentifier,omitempty" yaml:"sourceDbClusterIdentifier,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted bool `json:"storageEncrypted,omitempty" yaml:"storageEncrypted,omitempty"`

	// Name for an automatically created database on cluster creation.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
