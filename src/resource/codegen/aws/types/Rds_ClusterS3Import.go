package types

type Rds_ClusterS3Import struct {
	// Can be blank, but is the path to your backup
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// Role applied to load the data.
	IngestionRole string `json:"ingestionRole,omitempty" yaml:"ingestionRole,omitempty"`

	// Source engine for the backup
	SourceEngine string `json:"sourceEngine,omitempty" yaml:"sourceEngine,omitempty"`

	/*
	   Version of the source engine used to make the backup

	   This will not recreate the resource if the S3 object changes in some way. It's only used to initialize the database. This only works currently with the aurora engine. See AWS for currently supported engines and options. See [Aurora S3 Migration Docs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Migrating.ExtMySQL.html#AuroraMySQL.Migrating.ExtMySQL.S3).
	*/
	SourceEngineVersion string `json:"sourceEngineVersion,omitempty" yaml:"sourceEngineVersion,omitempty"`

	// Bucket name where your backup is stored
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
