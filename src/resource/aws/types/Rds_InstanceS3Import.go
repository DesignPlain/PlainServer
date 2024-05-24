package types

type Rds_InstanceS3Import struct {
	// Role applied to load the data.
	IngestionRole string `json:"ingestionRole,omitempty" yaml:"ingestionRole,omitempty"`

	// Source engine for the backup
	SourceEngine string `json:"sourceEngine,omitempty" yaml:"sourceEngine,omitempty"`

	/*
	   Version of the source engine used to make the backup

	   This will not recreate the resource if the S3 object changes in some way.  It's only used to initialize the database.
	*/
	SourceEngineVersion string `json:"sourceEngineVersion,omitempty" yaml:"sourceEngineVersion,omitempty"`

	// The bucket name where your backup is stored
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Can be blank, but is the path to your backup
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`
}
