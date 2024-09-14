package ec2

import types "libds/aws/types"

type Instance struct {
	// Number of the partition the instance is in. Valid only if the `aws.ec2.PlacementGroup` resource's `strategy` argument is set to `"partition"`.
	PlacementPartitionNumber int `json:"placementPartitionNumber,omitempty" yaml:"placementPartitionNumber,omitempty"`

	// List of secondary private IPv4 addresses to assign to the instance's primary network interface (eth0) in a VPC. Can only be assigned to the primary network interface (eth0) attached at instance creation, not a pre-existing network interface i.e., referenced in a `network_interface` block. Refer to the [Elastic network interfaces documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI) to see the maximum number of private IP addresses allowed per instance type.
	SecondaryPrivateIps []string `json:"secondaryPrivateIps,omitempty" yaml:"secondaryPrivateIps,omitempty"`

	// Whether to associate a public IP address with an instance in a VPC.
	AssociatePublicIpAddress bool `json:"associatePublicIpAddress,omitempty" yaml:"associatePublicIpAddress,omitempty"`

	// Sets the number of CPU cores for an instance. This option is only supported on creation of instance type that support CPU Options [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
	CpuCoreCount int `json:"cpuCoreCount,omitempty" yaml:"cpuCoreCount,omitempty"`

	// ARN of the host resource group in which to launch the instances. If you specify an ARN, omit the `tenancy` parameter or set it to `host`.
	HostResourceGroupArn string `json:"hostResourceGroupArn,omitempty" yaml:"hostResourceGroupArn,omitempty"`

	// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
	Monitoring bool `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	// AMI to use for the instance. Required unless `launch_template` is specified and the Launch Template specifes an AMI. If an AMI is specified in the Launch Template, setting `ami` will override the AMI specified in the Launch Template.
	Ami string `json:"ami,omitempty" yaml:"ami,omitempty"`

	/*
	   Describes an instance's Capacity Reservation targeting option. See Capacity Reservation Specification below for more details.

	   > --NOTE:-- Changing `cpu_core_count` and/or `cpu_threads_per_core` will cause the resource to be destroyed and re-created.
	*/
	CapacityReservationSpecification types.Ec2_InstanceCapacityReservationSpecification `json:"capacityReservationSpecification,omitempty" yaml:"capacityReservationSpecification,omitempty"`

	// If true, the launched EC2 instance will be EBS-optimized. Note that if this is not set on an instance type that is optimized by default then this will show as disabled but if the instance type is optimized by default then there is no need to set this and there is no effect to disabling it. See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
	EbsOptimized bool `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	// Map of tags to assign to the resource. Note that these tags apply to the instance and not block storage devices. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Describes the market (purchasing) option for the instances. See Market Options below for details on attributes.
	InstanceMarketOptions types.Ec2_InstanceInstanceMarketOptions `json:"instanceMarketOptions,omitempty" yaml:"instanceMarketOptions,omitempty"`

	// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" yaml:"ipv6Addresses,omitempty"`

	/*
	   List of security group names to associate with.

	   > --NOTE:-- If you are creating Instances in a VPC, use `vpc_security_group_ids` instead.
	*/
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// VPC Subnet ID to launch in.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions types.Ec2_InstanceCpuOptions `json:"cpuOptions,omitempty" yaml:"cpuOptions,omitempty"`

	// If set to 1, hyperthreading is disabled on the launched instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
	CpuThreadsPerCore int `json:"cpuThreadsPerCore,omitempty" yaml:"cpuThreadsPerCore,omitempty"`

	// One or more configuration blocks to customize Ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below for details. When accessing this as an attribute reference, it is a set of objects.
	EphemeralBlockDevices []types.Ec2_InstanceEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`

	// Shutdown behavior for the instance. Amazon defaults this to `stop` for EBS-backed instances and `terminate` for instance-store instances. Cannot be set on instance-store instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty" yaml:"instanceInitiatedShutdownBehavior,omitempty"`

	// When used in combination with `user_data` or `user_data_base64` will trigger a destroy and recreate when set to `true`. Defaults to `false` if not set.
	UserDataReplaceOnChange bool `json:"userDataReplaceOnChange,omitempty" yaml:"userDataReplaceOnChange,omitempty"`

	// If true, the launched EC2 instance will support hibernation.
	Hibernation bool `json:"hibernation,omitempty" yaml:"hibernation,omitempty"`

	// ID of a dedicated host that the instance will be assigned to. Use when an instance is to be launched on a specific dedicated host.
	HostId string `json:"hostId,omitempty" yaml:"hostId,omitempty"`

	/*
	   Map of tags to assign, at instance-creation time, to root and EBS volumes.

	   > --NOTE:-- Do not use `volume_tags` if you plan to manage block device tags outside the `aws.ec2.Instance` configuration, such as using `tags` in an `aws.ebs.Volume` resource attached via `aws.ec2.VolumeAttachment`. Doing so will result in resource cycling and inconsistent behavior.
	*/
	VolumeTags map[string]string `json:"volumeTags,omitempty" yaml:"volumeTags,omitempty"`

	// Configuration block for customizing the credit specification of the instance. See Credit Specification below for more details. This provider will only perform drift detection of its value when present in a configuration. Removing this configuration on existing instances will only stop managing it. It will not change the configuration back to the default for the instance type.
	CreditSpecification types.Ec2_InstanceCreditSpecification `json:"creditSpecification,omitempty" yaml:"creditSpecification,omitempty"`

	// If true, enables [EC2 Instance Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination).
	DisableApiTermination bool `json:"disableApiTermination,omitempty" yaml:"disableApiTermination,omitempty"`

	// One or more configuration blocks with additional EBS block devices to attach to the instance. Block device configurations only apply on resource creation. See Block Devices below for details on attributes and drift detection. When accessing this as an attribute reference, it is a set of objects.
	EbsBlockDevices []types.Ec2_InstanceEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions types.Ec2_InstanceEnclaveOptions `json:"enclaveOptions,omitempty" yaml:"enclaveOptions,omitempty"`

	// Number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" yaml:"ipv6AddressCount,omitempty"`

	// Private IP address to associate with the instance in a VPC.
	PrivateIp string `json:"privateIp,omitempty" yaml:"privateIp,omitempty"`

	// List of security group IDs to associate with.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// Tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of `dedicated` runs on single-tenant hardware. The `host` tenancy is not supported for the import-instance command. Valid values are `default`, `dedicated`, and `host`.
	Tenancy string `json:"tenancy,omitempty" yaml:"tenancy,omitempty"`

	// User data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead. Updates to this field will trigger a stop/start of the EC2 instance by default. If the `user_data_replace_on_change` is set then updates to this field will trigger a destroy and recreate.
	UserData string `json:"userData,omitempty" yaml:"userData,omitempty"`

	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData bool `json:"getPasswordData,omitempty" yaml:"getPasswordData,omitempty"`

	// IAM Instance Profile to launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" yaml:"iamInstanceProfile,omitempty"`

	// Key name of the Key Pair to use for the instance; which can be managed using the `aws.ec2.KeyPair` resource.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// Configuration block to customize details about the root block device of the instance. See Block Devices below for details. When accessing this as an attribute reference, it is a list containing one object.
	RootBlockDevice types.Ec2_InstanceRootBlockDevice `json:"rootBlockDevice,omitempty" yaml:"rootBlockDevice,omitempty"`

	// Placement Group to start the instance in.
	PlacementGroup string `json:"placementGroup,omitempty" yaml:"placementGroup,omitempty"`

	// Options for the instance hostname. The default values are inherited from the subnet. See Private DNS Name Options below for more details.
	PrivateDnsNameOptions types.Ec2_InstancePrivateDnsNameOptions `json:"privateDnsNameOptions,omitempty" yaml:"privateDnsNameOptions,omitempty"`

	// Controls if traffic is routed to the instance when the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
	SourceDestCheck bool `json:"sourceDestCheck,omitempty" yaml:"sourceDestCheck,omitempty"`

	// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption. Updates to this field will trigger a stop/start of the EC2 instance by default. If the `user_data_replace_on_change` is set then updates to this field will trigger a destroy and recreate.
	UserDataBase64 string `json:"userDataBase64,omitempty" yaml:"userDataBase64,omitempty"`

	// AZ to start the instance in.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Specifies a Launch Template to configure the instance. Parameters configured on this resource will override the corresponding parameters in the Launch Template. See Launch Template Specification below for more details.
	LaunchTemplate types.Ec2_InstanceLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`

	// Maintenance and recovery options for the instance. See Maintenance Options below for more details.
	MaintenanceOptions types.Ec2_InstanceMaintenanceOptions `json:"maintenanceOptions,omitempty" yaml:"maintenanceOptions,omitempty"`

	// Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
	NetworkInterfaces []types.Ec2_InstanceNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	// If true, enables [EC2 Instance Stop Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).
	DisableApiStop bool `json:"disableApiStop,omitempty" yaml:"disableApiStop,omitempty"`

	// Instance type to use for the instance. Required unless `launch_template` is specified and the Launch Template specifies an instance type. If an instance type is specified in the Launch Template, setting `instance_type` will override the instance type specified in the Launch Template. Updates to this field will trigger a stop/start of the EC2 instance.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Customize the metadata options of the instance. See Metadata Options below for more details.
	MetadataOptions types.Ec2_InstanceMetadataOptions `json:"metadataOptions,omitempty" yaml:"metadataOptions,omitempty"`
}
