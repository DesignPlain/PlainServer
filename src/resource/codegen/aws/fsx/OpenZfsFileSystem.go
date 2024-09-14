package fsx

import types "libds/aws/types"

type OpenZfsFileSystem struct {
	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToVolumes bool `json:"copyTagsToVolumes,omitempty" yaml:"copyTagsToVolumes,omitempty"`

	// A map of tags to apply to the file system's final backup.
	FinalBackupTags map[string]string `json:"finalBackupTags,omitempty" yaml:"finalBackupTags,omitempty"`

	/*
	   Throughput (MB/s) of the file system. Valid values depend on `deployment_type`. Must be one of `64`, `128`, `256`, `512`, `1024`, `2048`, `3072`, `4096` for `SINGLE_AZ_1`. Must be one of `160`, `320`, `640`, `1280`, `2560`, `3840`, `5120`, `7680`, `10240` for `SINGLE_AZ_2`.

	   The following arguments are optional:
	*/
	ThroughputCapacity int `json:"throughputCapacity,omitempty" yaml:"throughputCapacity,omitempty"`

	// The ID of the source backup to create the filesystem from.
	BackupId string `json:"backupId,omitempty" yaml:"backupId,omitempty"`

	// List of delete options, which at present supports only one value that specifies whether to delete all child volumes and snapshots when the file system is deleted. Valid values: `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS`.
	DeleteOptions []string `json:"deleteOptions,omitempty" yaml:"deleteOptions,omitempty"`

	// (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system will be created.
	EndpointIpAddressRange string `json:"endpointIpAddressRange,omitempty" yaml:"endpointIpAddressRange,omitempty"`

	// The configuration for the root volume of the file system. All other volumes are children or the root volume. See `root_volume_configuration` Block for details.
	RootVolumeConfiguration types.Fsx_OpenZfsFileSystemRootVolumeConfiguration `json:"rootVolumeConfiguration,omitempty" yaml:"rootVolumeConfiguration,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A map of tags to assign to the file system. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to backups. The default value is false.
	CopyTagsToBackups bool `json:"copyTagsToBackups,omitempty" yaml:"copyTagsToBackups,omitempty"`

	// The filesystem deployment type. Valid values: `SINGLE_AZ_1`, `SINGLE_AZ_2` and `MULTI_AZ_1`.
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`

	// The SSD IOPS configuration for the Amazon FSx for OpenZFS file system. See `disk_iops_configuration` Block for details.
	DiskIopsConfiguration types.Fsx_OpenZfsFileSystemDiskIopsConfiguration `json:"diskIopsConfiguration,omitempty" yaml:"diskIopsConfiguration,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup bool `json:"skipFinalBackup,omitempty" yaml:"skipFinalBackup,omitempty"`

	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime string `json:"weeklyMaintenanceStartTime,omitempty" yaml:"weeklyMaintenanceStartTime,omitempty"`

	// The filesystem storage type. Only `SSD` is supported.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// A list of IDs for the subnets that the file system will be accessible from.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays int `json:"automaticBackupRetentionDays,omitempty" yaml:"automaticBackupRetentionDays,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automatic_backup_retention_days` to be set.
	DailyAutomaticBackupStartTime string `json:"dailyAutomaticBackupStartTime,omitempty" yaml:"dailyAutomaticBackupStartTime,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// (Multi-AZ only) Required when `deployment_type` is set to `MULTI_AZ_1`. This specifies the subnet in which you want the preferred file server to be located.
	PreferredSubnetId string `json:"preferredSubnetId,omitempty" yaml:"preferredSubnetId,omitempty"`

	// (Multi-AZ only) Specifies the route tables in which Amazon FSx creates the rules for routing traffic to the correct file server. You should specify all virtual private cloud (VPC) route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds []string `json:"routeTableIds,omitempty" yaml:"routeTableIds,omitempty"`

	// The storage capacity (GiB) of the file system. Valid values between `64` and `524288`.
	StorageCapacity int `json:"storageCapacity,omitempty" yaml:"storageCapacity,omitempty"`
}
