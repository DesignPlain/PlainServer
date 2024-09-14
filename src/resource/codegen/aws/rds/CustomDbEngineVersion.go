package rds

type CustomDbEngineVersion struct {
	// The prefix for the Amazon S3 bucket that contains the database installation files.
	DatabaseInstallationFilesS3Prefix string `json:"databaseInstallationFilesS3Prefix,omitempty" yaml:"databaseInstallationFilesS3Prefix,omitempty"`

	// The description of the CEV.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ARN of the AWS KMS key that is used to encrypt the database installation files. Required for RDS Custom for Oracle.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The manifest file, in JSON format, that contains the list of database installation files. Conflicts with `filename`.
	Manifest string `json:"manifest,omitempty" yaml:"manifest,omitempty"`

	// The ID of the AMI to create the CEV from. Required for RDS Custom for SQL Server. For RDS Custom for Oracle, you can specify an AMI ID that was used in a different Oracle CEV.
	SourceImageId string `json:"sourceImageId,omitempty" yaml:"sourceImageId,omitempty"`

	// The name of the Amazon S3 bucket that contains the database installation files.
	DatabaseInstallationFilesS3BucketName string `json:"databaseInstallationFilesS3BucketName,omitempty" yaml:"databaseInstallationFilesS3BucketName,omitempty"`

	// The version of the database engine.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// The name of the manifest file within the local filesystem. Conflicts with `manifest`.
	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the manifest source specified with `filename`. The usual way to set this is filebase64sha256("manifest.json") where "manifest.json" is the local filename of the manifest source.
	ManifestHash string `json:"manifestHash,omitempty" yaml:"manifestHash,omitempty"`

	// The status of the CEV. Valid values are `available`, `inactive`, `inactive-except-restore`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the database engine. Valid values are `custom-oracle-`, `custom-sqlserver-`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`
}
