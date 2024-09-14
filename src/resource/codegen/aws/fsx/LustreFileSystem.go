package fsx

import types "libds/aws/types"

type LustreFileSystem struct {
	// A map of tags to assign to the file system. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see [Auto Import Data Repo](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) for more details. Only supported on `PERSISTENT_1` deployment types.
	AutoImportPolicy string `json:"autoImportPolicy,omitempty" yaml:"autoImportPolicy,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for `PERSISTENT_1` and `PERSISTENT_2` deployment_type. The default value is false.
	CopyTagsToBackups bool `json:"copyTagsToBackups,omitempty" yaml:"copyTagsToBackups,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for `PERSISTENT_1` and `PERSISTENT_2` deployment_type. Requires `automatic_backup_retention_days` to be set.
	DailyAutomaticBackupStartTime string `json:"dailyAutomaticBackupStartTime,omitempty" yaml:"dailyAutomaticBackupStartTime,omitempty"`

	// The type of drive cache used by `PERSISTENT_1` filesystems that are provisioned with `HDD` storage_type. Required for `HDD` storage_type, set to either `READ` or `NONE`.
	DriveCacheType string `json:"driveCacheType,omitempty" yaml:"driveCacheType,omitempty"`

	// The Lustre metadata configuration used when creating an Amazon FSx for Lustre file system. This can be used to specify a user provisioned metadata scale. This is only supported when `deployment_type` is set to `PERSISTENT_2`. See `metadata_configuration` Block for details.
	MetadataConfiguration types.Fsx_LustreFileSystemMetadataConfiguration `json:"metadataConfiguration,omitempty" yaml:"metadataConfiguration,omitempty"`

	// Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` and `PERSISTENT_2` deployment_type. Valid values for `PERSISTENT_1` deployment_type and `SSD` storage_type are 50, 100, 200. Valid values for `PERSISTENT_1` deployment_type and `HDD` storage_type are 12, 40. Valid values for `PERSISTENT_2` deployment_type and ` SSD` storage_type are 125, 250, 500, 1000.
	PerUnitStorageThroughput int `json:"perUnitStorageThroughput,omitempty" yaml:"perUnitStorageThroughput,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, applicable for `PERSISTENT_1` and `PERSISTENT_2` deployment_type. Defaults to an AWS managed KMS Key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The storage capacity (GiB) of the file system. Minimum of `1200`. See more details at [Allowed values for Fsx storage capacity](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystem.html#FSx-CreateFileSystem-request-StorageCapacity). Update is allowed only for `SCRATCH_2`, `PERSISTENT_1` and `PERSISTENT_2` deployment types, See more details at [Fsx Storage Capacity Update](https://docs.aws.amazon.com/fsx/latest/APIReference/API_UpdateFileSystem.html#FSx-UpdateFileSystem-request-StorageCapacity). Required when not creating filesystem for a backup.
	StorageCapacity int `json:"storageCapacity,omitempty" yaml:"storageCapacity,omitempty"`

	// The filesystem storage type. Either `SSD` or `HDD`, defaults to `SSD`. `HDD` is only supported on `PERSISTENT_1` deployment types.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`, `PERSISTENT_2`.
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`

	/*
	   A map of tags to apply to the file system's final backup.

	   --Note:-- If the filesystem uses a Scratch deployment type, final backup during delete will always be skipped and this argument will not be used even when set.
	*/
	FinalBackupTags map[string]string `json:"finalBackupTags,omitempty" yaml:"finalBackupTags,omitempty"`

	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`. Only supported on `PERSISTENT_1` deployment types.
	ImportPath string `json:"importPath,omitempty" yaml:"importPath,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`. Only supported on `PERSISTENT_1` deployment types.
	ImportedFileChunkSize int `json:"importedFileChunkSize,omitempty" yaml:"importedFileChunkSize,omitempty"`

	/*
	   A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.

	   The following arguments are optional:
	*/
	SubnetIds string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime string `json:"weeklyMaintenanceStartTime,omitempty" yaml:"weeklyMaintenanceStartTime,omitempty"`

	// The Lustre root squash configuration used when creating an Amazon FSx for Lustre file system. When enabled, root squash restricts root-level access from clients that try to access your file system as a root user. See `root_squash_configuration` Block for details.
	RootSquashConfiguration types.Fsx_LustreFileSystemRootSquashConfiguration `json:"rootSquashConfiguration,omitempty" yaml:"rootSquashConfiguration,omitempty"`

	/*
	   When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `true`.

	   --Note:-- If the filesystem uses a Scratch deployment type, final backup during delete will always be skipped and this argument will not be used even when set.
	*/
	SkipFinalBackup bool `json:"skipFinalBackup,omitempty" yaml:"skipFinalBackup,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for `PERSISTENT_1` and `PERSISTENT_2` deployment_type.
	AutomaticBackupRetentionDays int `json:"automaticBackupRetentionDays,omitempty" yaml:"automaticBackupRetentionDays,omitempty"`

	// The ID of the source backup to create the filesystem from.
	BackupId string `json:"backupId,omitempty" yaml:"backupId,omitempty"`

	// Sets the data compression configuration for the file system. Valid values are `LZ4` and `NONE`. Default value is `NONE`. Unsetting this value reverts the compression type back to `NONE`.
	DataCompressionType string `json:"dataCompressionType,omitempty" yaml:"dataCompressionType,omitempty"`

	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`. Only supported on `PERSISTENT_1` deployment types.
	ExportPath string `json:"exportPath,omitempty" yaml:"exportPath,omitempty"`

	// Sets the Lustre version for the file system that you're creating. Valid values are 2.10 for `SCRATCH_1`, `SCRATCH_2` and `PERSISTENT_1` deployment types. Valid values for 2.12 include all deployment types.
	FileSystemTypeVersion string `json:"fileSystemTypeVersion,omitempty" yaml:"fileSystemTypeVersion,omitempty"`

	// The Lustre logging configuration used when creating an Amazon FSx for Lustre file system. When logging is enabled, Lustre logs error and warning events for data repositories associated with your file system to Amazon CloudWatch Logs. See `log_configuration` Block for details.
	LogConfiguration types.Fsx_LustreFileSystemLogConfiguration `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`
}
