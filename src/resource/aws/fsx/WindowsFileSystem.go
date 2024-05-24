package fsx

import types "DesignSphere_Server/src/resource/aws/types"

type WindowsFileSystem struct {
	// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system. See Audit Log Configuration below.
	AuditLogConfiguration types.Fsx_WindowsFileSystemAuditLogConfiguration `json:"auditLogConfiguration,omitempty" yaml:"auditLogConfiguration,omitempty"`

	// The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
	AutomaticBackupRetentionDays int `json:"automaticBackupRetentionDays,omitempty" yaml:"automaticBackupRetentionDays,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
	PreferredSubnetId string `json:"preferredSubnetId,omitempty" yaml:"preferredSubnetId,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime string `json:"weeklyMaintenanceStartTime,omitempty" yaml:"weeklyMaintenanceStartTime,omitempty"`

	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `self_managed_active_directory`.
	ActiveDirectoryId string `json:"activeDirectoryId,omitempty" yaml:"activeDirectoryId,omitempty"`

	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
	CopyTagsToBackups bool `json:"copyTagsToBackups,omitempty" yaml:"copyTagsToBackups,omitempty"`

	// The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime string `json:"dailyAutomaticBackupStartTime,omitempty" yaml:"dailyAutomaticBackupStartTime,omitempty"`

	// Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup bool `json:"skipFinalBackup,omitempty" yaml:"skipFinalBackup,omitempty"`

	// An array DNS alias names that you want to associate with the Amazon FSx file system.  For more information, see [Working with DNS Aliases](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-dns-aliases.html)
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty"`

	// The ID of the source backup to create the filesystem from.
	BackupId string `json:"backupId,omitempty" yaml:"backupId,omitempty"`

	// The SSD IOPS configuration for the Amazon FSx for Windows File Server file system. See Disk Iops Configuration below.
	DiskIopsConfiguration types.Fsx_WindowsFileSystemDiskIopsConfiguration `json:"diskIopsConfiguration,omitempty" yaml:"diskIopsConfiguration,omitempty"`

	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000. Required when not creating filesystem for a backup.
	StorageCapacity int `json:"storageCapacity,omitempty" yaml:"storageCapacity,omitempty"`

	// Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// A map of tags to assign to the file system. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.

	   The following arguments are optional:
	*/
	ThroughputCapacity int `json:"throughputCapacity,omitempty" yaml:"throughputCapacity,omitempty"`

	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `active_directory_id`. See Self-Managed Active Directory below.
	SelfManagedActiveDirectory types.Fsx_WindowsFileSystemSelfManagedActiveDirectory `json:"selfManagedActiveDirectory,omitempty" yaml:"selfManagedActiveDirectory,omitempty"`

	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deployment_type` to `MULTI_AZ_1`.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
