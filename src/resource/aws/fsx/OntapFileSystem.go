package fsx

import types "DesignSphere_Server/src/resource/aws/types"

type OntapFileSystem struct {
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime string `json:"weeklyMaintenanceStartTime,omitempty" yaml:"weeklyMaintenanceStartTime,omitempty"`

	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.- range.
	EndpointIpAddressRange string `json:"endpointIpAddressRange,omitempty" yaml:"endpointIpAddressRange,omitempty"`

	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetId string `json:"preferredSubnetId,omitempty" yaml:"preferredSubnetId,omitempty"`

	// The filesystem storage type. defaults to `SSD`.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// A list of IDs for the subnets that the file system will be accessible from. Up to 2 subnets can be provided.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automatic_backup_retention_days` to be set.
	DailyAutomaticBackupStartTime string `json:"dailyAutomaticBackupStartTime,omitempty" yaml:"dailyAutomaticBackupStartTime,omitempty"`

	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FsxAdminPassword string `json:"fsxAdminPassword,omitempty" yaml:"fsxAdminPassword,omitempty"`

	// The number of ha_pairs to deploy for the file system. Valid values are 1 through 6. Recommend only using this parameter for 2 or more ha pairs.
	HaPairs int `json:"haPairs,omitempty" yaml:"haPairs,omitempty"`

	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds []string `json:"routeTableIds,omitempty" yaml:"routeTableIds,omitempty"`

	// A map of tags to assign to the file system. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays int `json:"automaticBackupRetentionDays,omitempty" yaml:"automaticBackupRetentionDays,omitempty"`

	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration below.
	DiskIopsConfiguration types.Fsx_OntapFileSystemDiskIopsConfiguration `json:"diskIopsConfiguration,omitempty" yaml:"diskIopsConfiguration,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
	StorageCapacity int `json:"storageCapacity,omitempty" yaml:"storageCapacity,omitempty"`

	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, `2048`, and `4096`. This parameter should only be used when specifying not using the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	ThroughputCapacity int `json:"throughputCapacity,omitempty" yaml:"throughputCapacity,omitempty"`

	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `3072`,`6144`. This parameter should only be used when specifying the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	ThroughputCapacityPerHaPair int `json:"throughputCapacityPerHaPair,omitempty" yaml:"throughputCapacityPerHaPair,omitempty"`

	// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`
}