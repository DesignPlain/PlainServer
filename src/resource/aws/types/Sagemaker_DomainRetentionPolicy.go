package types

type Sagemaker_DomainRetentionPolicy struct {
	// The retention policy for data stored on an Amazon Elastic File System (EFS) volume. Valid values are `Retain` or `Delete`.  Default value is `Retain`.
	HomeEfsFileSystem string `json:"homeEfsFileSystem,omitempty" yaml:"homeEfsFileSystem,omitempty"`
}
