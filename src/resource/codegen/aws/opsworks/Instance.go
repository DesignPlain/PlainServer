package opsworks

import types "libds/aws/types"

type Instance struct {
	// Configuration block for ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below.
	EphemeralBlockDevices []types.Opsworks_InstanceEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`

	// List of the layers the instance will belong to.
	LayerIds []string `json:"layerIds,omitempty" yaml:"layerIds,omitempty"`

	/*
	   Identifier of the stack the instance will belong to.

	   The following arguments are optional:
	*/
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`

	// Machine architecture for created instances.  Valid values are `x86_64` or `i386`. The default is `x86_64`.
	Architecture string `json:"architecture,omitempty" yaml:"architecture,omitempty"`

	// Controls where to install OS and package updates when the instance boots.  Default is `true`.
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" yaml:"installUpdatesOnBoot,omitempty"`

	// Name of the SSH keypair that instances will have by default.
	SshKeyName string `json:"sshKeyName,omitempty" yaml:"sshKeyName,omitempty"`

	// Instance tenancy to use. Valid values are `default`, `dedicated` or `host`.
	Tenancy string `json:"tenancy,omitempty" yaml:"tenancy,omitempty"`

	// Name of the availability zone where instances will be created by default.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Instance's host name.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// Configuration block for the root block device of the instance. See Block Devices below.
	RootBlockDevices []types.Opsworks_InstanceRootBlockDevice `json:"rootBlockDevices,omitempty" yaml:"rootBlockDevices,omitempty"`

	// Name of the type of root device instances will have by default. Valid values are `ebs` or `instance-store`.
	RootDeviceType string `json:"rootDeviceType,omitempty" yaml:"rootDeviceType,omitempty"`

	// Subnet ID to attach to.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// OpsWorks agent to install. Default is `INHERIT`.
	AgentVersion string `json:"agentVersion,omitempty" yaml:"agentVersion,omitempty"`

	// Type of instance to start.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Instance status. Will be one of `booting`, `connection_lost`, `online`, `pending`, `rebooting`, `requested`, `running_setup`, `setup_failed`, `shutting_down`, `start_failed`, `stop_failed`, `stopped`, `stopping`, `terminated`, or `terminating`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Whether to delete the Elastic IP on deletion.
	DeleteEip bool `json:"deleteEip,omitempty" yaml:"deleteEip,omitempty"`

	// Whether the launched EC2 instance will be EBS-optimized.
	EbsOptimized bool `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	// ECS cluster's ARN for container instances.
	EcsClusterArn string `json:"ecsClusterArn,omitempty" yaml:"ecsClusterArn,omitempty"`

	// For registered instances, infrastructure class: ec2 or on-premises.
	InfrastructureClass string `json:"infrastructureClass,omitempty" yaml:"infrastructureClass,omitempty"`

	// Name of operating system that will be installed.
	Os string `json:"os,omitempty" yaml:"os,omitempty"`

	// Associated security groups.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Whether to delete EBS volume on deletion. Default is `true`.
	DeleteEbs bool `json:"deleteEbs,omitempty" yaml:"deleteEbs,omitempty"`

	// Creates load-based or time-based instances.  Valid values are `load`, `timer`.
	AutoScalingType string `json:"autoScalingType,omitempty" yaml:"autoScalingType,omitempty"`

	// Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
	EbsBlockDevices []types.Opsworks_InstanceEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	// AMI to use for the instance.  If an AMI is specified, `os` must be `Custom`.
	AmiId string `json:"amiId,omitempty" yaml:"amiId,omitempty"`

	// Instance Elastic IP address.
	ElasticIp string `json:"elasticIp,omitempty" yaml:"elasticIp,omitempty"`

	// ARN of the instance's IAM profile.
	InstanceProfileArn string `json:"instanceProfileArn,omitempty" yaml:"instanceProfileArn,omitempty"`

	// Desired state of the instance. Valid values are `running` or `stopped`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Keyword to choose what virtualization mode created instances will use. Valid values are `paravirtual` or `hvm`.
	VirtualizationType string `json:"virtualizationType,omitempty" yaml:"virtualizationType,omitempty"`

	// Time that the instance was created.
	CreatedAt string `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
}
