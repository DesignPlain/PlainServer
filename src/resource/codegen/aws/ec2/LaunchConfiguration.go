package ec2

import types "libds/aws/types"

type LaunchConfiguration struct {
	// The EC2 image ID to launch.
	ImageId string `json:"imageId,omitempty" yaml:"imageId,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// A list of associated security group IDS.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized bool `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring bool `json:"enableMonitoring,omitempty" yaml:"enableMonitoring,omitempty"`

	// Customize Ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices []types.Ec2_LaunchConfigurationEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`

	// The key name that should be used for the instance.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress bool `json:"associatePublicIpAddress,omitempty" yaml:"associatePublicIpAddress,omitempty"`

	// Additional EBS block devices to attach to the instance. See Block Devices below for details.
	EbsBlockDevices []types.Ec2_LaunchConfigurationEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	// The name attribute of the IAM instance profile to associate with launched instances.
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" yaml:"iamInstanceProfile,omitempty"`

	// The tenancy of the instance. Valid values are `default` or `dedicated`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html) for more details.
	PlacementTenancy string `json:"placementTenancy,omitempty" yaml:"placementTenancy,omitempty"`

	// Customize details about the root block device of the instance. See Block Devices below for details.
	RootBlockDevice types.Ec2_LaunchConfigurationRootBlockDevice `json:"rootBlockDevice,omitempty" yaml:"rootBlockDevice,omitempty"`

	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
	UserData string `json:"userData,omitempty" yaml:"userData,omitempty"`

	/*
	   The size of instance to launch.

	   The following arguments are optional:
	*/
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The metadata options for the instance.
	MetadataOptions types.Ec2_LaunchConfigurationMetadataOptions `json:"metadataOptions,omitempty" yaml:"metadataOptions,omitempty"`

	// The name of the launch configuration. If you leave this blank, this provider will auto-generate a unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The maximum price to use for reserving spot instances.
	SpotPrice string `json:"spotPrice,omitempty" yaml:"spotPrice,omitempty"`

	// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 string `json:"userDataBase64,omitempty" yaml:"userDataBase64,omitempty"`
}
