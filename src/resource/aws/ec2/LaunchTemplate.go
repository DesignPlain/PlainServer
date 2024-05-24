package ec2

import types "DesignSphere_Server/src/resource/aws/types"

type LaunchTemplate struct {
	// The name of the launch template. If you leave this blank, the provider will auto-generate a unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The options for the instance hostname. The default values are inherited from the subnet. See Private DNS Name Options below for more details.
	PrivateDnsNameOptions types.Ec2_LaunchTemplatePrivateDnsNameOptions `json:"privateDnsNameOptions,omitempty" yaml:"privateDnsNameOptions,omitempty"`

	// Description of the launch template.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   If `true`, enables [EC2 Instance
	   Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	*/
	DisableApiTermination bool `json:"disableApiTermination,omitempty" yaml:"disableApiTermination,omitempty"`

	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized string `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	// The AMI from which to launch the instance.
	ImageId string `json:"imageId,omitempty" yaml:"imageId,omitempty"`

	// The kernel ID.
	KernelId string `json:"kernelId,omitempty" yaml:"kernelId,omitempty"`

	// The maintenance options for the instance. See Maintenance Options below for more details.
	MaintenanceOptions types.Ec2_LaunchTemplateMaintenanceOptions `json:"maintenanceOptions,omitempty" yaml:"maintenanceOptions,omitempty"`

	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications []types.Ec2_LaunchTemplateLicenseSpecification `json:"licenseSpecifications,omitempty" yaml:"licenseSpecifications,omitempty"`

	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring types.Ec2_LaunchTemplateMonitoring `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   Customize network interfaces to be attached at instance boot time. See Network
	   Interfaces below for more details.
	*/
	NetworkInterfaces []types.Ec2_LaunchTemplateNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	// The base64-encoded user data to provide when launching the instance.
	UserData string `json:"userData,omitempty" yaml:"userData,omitempty"`

	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification types.Ec2_LaunchTemplateCapacityReservationSpecification `json:"capacityReservationSpecification,omitempty" yaml:"capacityReservationSpecification,omitempty"`

	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions types.Ec2_LaunchTemplateEnclaveOptions `json:"enclaveOptions,omitempty" yaml:"enclaveOptions,omitempty"`

	/*
	   Shutdown behavior for the instance. Can be `stop` or `terminate`.
	   (Default: `stop`).
	*/
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty" yaml:"instanceInitiatedShutdownBehavior,omitempty"`

	/*
	   The market (purchasing) option for the instance. See Market Options
	   below for details.
	*/
	InstanceMarketOptions types.Ec2_LaunchTemplateInstanceMarketOptions `json:"instanceMarketOptions,omitempty" yaml:"instanceMarketOptions,omitempty"`

	// The attribute requirements for the type of instance. If present then `instance_type` cannot be present.
	InstanceRequirements types.Ec2_LaunchTemplateInstanceRequirements `json:"instanceRequirements,omitempty" yaml:"instanceRequirements,omitempty"`

	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications []types.Ec2_LaunchTemplateTagSpecification `json:"tagSpecifications,omitempty" yaml:"tagSpecifications,omitempty"`

	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions types.Ec2_LaunchTemplateCpuOptions `json:"cpuOptions,omitempty" yaml:"cpuOptions,omitempty"`

	/*
	   Customize the credit specification of the instance. See Credit
	   Specification below for more details.
	*/
	CreditSpecification types.Ec2_LaunchTemplateCreditSpecification `json:"creditSpecification,omitempty" yaml:"creditSpecification,omitempty"`

	// Default Version of the launch template.
	DefaultVersion int `json:"defaultVersion,omitempty" yaml:"defaultVersion,omitempty"`

	// If true, enables [EC2 Instance Stop Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).
	DisableApiStop bool `json:"disableApiStop,omitempty" yaml:"disableApiStop,omitempty"`

	// A map of tags to assign to the launch template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether to update Default Version each update. Conflicts with `default_version`.
	UpdateDefaultVersion bool `json:"updateDefaultVersion,omitempty" yaml:"updateDefaultVersion,omitempty"`

	/*
	   Specify volumes to attach to the instance besides the volumes specified by the AMI.
	   See Block Devices below for details.
	*/
	BlockDeviceMappings []types.Ec2_LaunchTemplateBlockDeviceMapping `json:"blockDeviceMappings,omitempty" yaml:"blockDeviceMappings,omitempty"`

	/*
	   The IAM Instance Profile to launch the instance with. See Instance Profile
	   below for more details.
	*/
	IamInstanceProfile types.Ec2_LaunchTemplateIamInstanceProfile `json:"iamInstanceProfile,omitempty" yaml:"iamInstanceProfile,omitempty"`

	// Customize the metadata options for the instance. See Metadata Options below for more details.
	MetadataOptions types.Ec2_LaunchTemplateMetadataOptions `json:"metadataOptions,omitempty" yaml:"metadataOptions,omitempty"`

	/*
	   A list of security group names to associate with. If you are creating Instances in a VPC, use
	   `vpc_security_group_ids` instead.
	*/
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" yaml:"securityGroupNames,omitempty"`

	// A list of security group IDs to associate with. Conflicts with `network_interfaces.security_groups`
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator types.Ec2_LaunchTemplateElasticInferenceAccelerator `json:"elasticInferenceAccelerator,omitempty" yaml:"elasticInferenceAccelerator,omitempty"`

	// The hibernation options for the instance. See Hibernation Options below for more details.
	HibernationOptions types.Ec2_LaunchTemplateHibernationOptions `json:"hibernationOptions,omitempty" yaml:"hibernationOptions,omitempty"`

	// The key name to use for the instance.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// The placement of the instance. See Placement below for more details.
	Placement types.Ec2_LaunchTemplatePlacement `json:"placement,omitempty" yaml:"placement,omitempty"`

	// The ID of the RAM disk.
	RamDiskId string `json:"ramDiskId,omitempty" yaml:"ramDiskId,omitempty"`

	/*
	   The elastic GPU to attach to the instance. See Elastic GPU
	   below for more details.
	*/
	ElasticGpuSpecifications []types.Ec2_LaunchTemplateElasticGpuSpecification `json:"elasticGpuSpecifications,omitempty" yaml:"elasticGpuSpecifications,omitempty"`

	// The type of the instance. If present then `instance_requirements` cannot be present.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
}
