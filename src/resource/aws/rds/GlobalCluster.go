package rds

type GlobalCluster struct {
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier string `json:"sourceDbClusterIdentifier,omitempty" yaml:"sourceDbClusterIdentifier,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted bool `json:"storageEncrypted,omitempty" yaml:"storageEncrypted,omitempty"`

	// Name for an automatically created database on cluster creation.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `source_db_cluster_identifier`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Engine version of the Aurora global database. The `engine`, `engine_version`, and `instance_class` (on the `aws.rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. --NOTE:-- To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignore_changes` for `engine_version` meta argument on the associated `aws.rds.Cluster` resource as shown above in Upgrading Engine Versions example.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `source_db_cluster_identifier`.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Global cluster identifier.
	GlobalClusterIdentifier string `json:"globalClusterIdentifier,omitempty" yaml:"globalClusterIdentifier,omitempty"`
}
