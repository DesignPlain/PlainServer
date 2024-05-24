package fsx

import types "DesignSphere_Server/src/resource/aws/types"

type FileCache struct {
	// A boolean flag indicating whether tags for the cache should be copied to data repository associations. This value defaults to false.
	CopyTagsToDataRepositoryAssociations bool `json:"copyTagsToDataRepositoryAssociations,omitempty" yaml:"copyTagsToDataRepositoryAssociations,omitempty"`

	// The type of cache that you're creating. The only supported value is `LUSTRE`.
	FileCacheType string `json:"fileCacheType,omitempty" yaml:"fileCacheType,omitempty"`

	// The version for the type of cache that you're creating. The only supported value is `2.12`.
	FileCacheTypeVersion string `json:"fileCacheTypeVersion,omitempty" yaml:"fileCacheTypeVersion,omitempty"`

	// See the `lustre_configuration` block. Required when `file_cache_type` is `LUSTRE`.
	LustreConfigurations []types.Fsx_FileCacheLustreConfiguration `json:"lustreConfigurations,omitempty" yaml:"lustreConfigurations,omitempty"`

	/*
	   A list of subnet IDs that the cache will be accessible from. You can specify only one subnet ID.

	   The following arguments are optional:
	*/
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   See the `data_repository_association` configuration block. Max of 8.
	   A list of up to 8 configurations for data repository associations (DRAs) to be created during the cache creation. The DRAs link the cache to either an Amazon S3 data repository or a Network File System (NFS) data repository that supports the NFSv3 protocol. The DRA configurations must meet the following requirements: 1) All configurations on the list must be of the same data repository type, either all S3 or all NFS. A cache can't link to different data repository types at the same time. 2) An NFS DRA must link to an NFS file system that supports the NFSv3 protocol. DRA automatic import and automatic export is not supported.
	*/
	DataRepositoryAssociations []types.Fsx_FileCacheDataRepositoryAssociation `json:"dataRepositoryAssociations,omitempty" yaml:"dataRepositoryAssociations,omitempty"`

	// Specifies the ID of the AWS Key Management Service (AWS KMS) key to use for encrypting data on an Amazon File Cache. If a KmsKeyId isn't specified, the Amazon FSx-managed AWS KMS key for your account is used.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A list of IDs specifying the security groups to apply to all network interfaces created for Amazon File Cache access.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The storage capacity of the cache in gibibytes (GiB). Valid values are `1200` GiB, `2400` GiB, and increments of `2400` GiB.
	StorageCapacity int `json:"storageCapacity,omitempty" yaml:"storageCapacity,omitempty"`
}
