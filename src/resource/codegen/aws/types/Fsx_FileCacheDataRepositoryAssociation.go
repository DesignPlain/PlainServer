package types

type Fsx_FileCacheDataRepositoryAssociation struct {
	// The path to the S3 or NFS data repository that links to the cache.
	DataRepositoryPath string `json:"dataRepositoryPath,omitempty" yaml:"dataRepositoryPath,omitempty"`

	// A path on the cache that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath. The leading forward slash in the name is required. Two data repository associations cannot have overlapping cache paths. For example, if a data repository is associated with cache path /ns1/, then you cannot link another data repository with cache path /ns1/ns2. This path specifies where in your cache files will be exported from. This cache directory can be linked to only one data repository, and no data repository other can be linked to the directory. Note: The cache path can only be set to root (/) on an NFS DRA when DataRepositorySubdirectories is specified. If you specify root (/) as the cache path, you can create only one DRA on the cache. The cache path cannot be set to root (/) for an S3 DRA.
	FileCachePath string `json:"fileCachePath,omitempty" yaml:"fileCachePath,omitempty"`

	//
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	//
	FileSystemPath string `json:"fileSystemPath,omitempty" yaml:"fileSystemPath,omitempty"`

	// (Optional) See the `nfs` configuration block.
	Nfs []Fsx_FileCacheDataRepositoryAssociationNf `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	//
	AssociationId string `json:"associationId,omitempty" yaml:"associationId,omitempty"`

	// A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
	DataRepositorySubdirectories []string `json:"dataRepositorySubdirectories,omitempty" yaml:"dataRepositorySubdirectories,omitempty"`

	// The system-generated, unique ID of the cache.
	FileCacheId string `json:"fileCacheId,omitempty" yaml:"fileCacheId,omitempty"`

	//
	ImportedFileChunkSize int `json:"importedFileChunkSize,omitempty" yaml:"importedFileChunkSize,omitempty"`

	//
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
