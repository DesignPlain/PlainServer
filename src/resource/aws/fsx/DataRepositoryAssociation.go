package fsx

import types "DesignSphere_Server/src/resource/aws/types"

type DataRepositoryAssociation struct {
	// A path on the file system that points to a high-level directory (such as `/ns1/`) or subdirectory (such as `/ns1/subdir/`) that will be mapped 1-1 with `data_repository_path`. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path `/ns1/`, then you cannot link another data repository with file system path `/ns1/ns2`. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	FileSystemPath string `json:"fileSystemPath,omitempty" yaml:"fileSystemPath,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	ImportedFileChunkSize int `json:"importedFileChunkSize,omitempty" yaml:"importedFileChunkSize,omitempty"`

	/*
	   See the `s3` configuration block. Max of 1.
	   The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
	*/
	S3 types.Fsx_DataRepositoryAssociationS3 `json:"s3,omitempty" yaml:"s3,omitempty"`

	// A map of tags to assign to the data repository association. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to `false`.
	BatchImportMetaDataOnCreate bool `json:"batchImportMetaDataOnCreate,omitempty" yaml:"batchImportMetaDataOnCreate,omitempty"`

	// The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
	DataRepositoryPath string `json:"dataRepositoryPath,omitempty" yaml:"dataRepositoryPath,omitempty"`

	// Set to true to delete files from the file system upon deleting this data repository association. Defaults to `false`.
	DeleteDataInFilesystem bool `json:"deleteDataInFilesystem,omitempty" yaml:"deleteDataInFilesystem,omitempty"`

	// The ID of the Amazon FSx file system to on which to create a data repository association.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`
}
