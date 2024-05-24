package opsworks

import types "DesignSphere_Server/src/resource/aws/types"

type Stack struct {
	// The ARN of an IAM Instance Profile that created instances will have by default.
	DefaultInstanceProfileArn string `json:"defaultInstanceProfileArn,omitempty" yaml:"defaultInstanceProfileArn,omitempty"`

	// The name of the stack.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A map of tags to assign to the resource.
	   If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// If `manage_berkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion string `json:"berkshelfVersion,omitempty" yaml:"berkshelfVersion,omitempty"`

	// When `use_custom_cookbooks` is set, provide this sub-object as described below.
	CustomCookbooksSources []types.Opsworks_StackCustomCookbooksSource `json:"customCookbooksSources,omitempty" yaml:"customCookbooksSources,omitempty"`

	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType string `json:"defaultRootDeviceType,omitempty" yaml:"defaultRootDeviceType,omitempty"`

	// Boolean value controlling whether the standard OpsWorks security groups apply to created instances.
	UseOpsworksSecurityGroups bool `json:"useOpsworksSecurityGroups,omitempty" yaml:"useOpsworksSecurityGroups,omitempty"`

	// The name of the region where the stack will exist.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" yaml:"serviceRoleArn,omitempty"`

	// Boolean value controlling whether the custom cookbook settings are enabled.
	UseCustomCookbooks bool `json:"useCustomCookbooks,omitempty" yaml:"useCustomCookbooks,omitempty"`

	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion string `json:"agentVersion,omitempty" yaml:"agentVersion,omitempty"`

	// Color to paint next to the stack's resources in the OpsWorks console.
	Color string `json:"color,omitempty" yaml:"color,omitempty"`

	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName string `json:"configurationManagerName,omitempty" yaml:"configurationManagerName,omitempty"`

	// Name of OS that will be installed on instances by default.
	DefaultOs string `json:"defaultOs,omitempty" yaml:"defaultOs,omitempty"`

	/*
	   ID of the subnet in which instances will be created by default.
	   Required if `vpc_id` is set to a VPC other than the default VPC, and forbidden if it isn't.
	*/
	DefaultSubnetId string `json:"defaultSubnetId,omitempty" yaml:"defaultSubnetId,omitempty"`

	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf bool `json:"manageBerkshelf,omitempty" yaml:"manageBerkshelf,omitempty"`

	/*
	   ID of the VPC that this stack belongs to.
	   Defaults to the region's default VPC.
	*/
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion string `json:"configurationManagerVersion,omitempty" yaml:"configurationManagerVersion,omitempty"`

	// Custom JSON attributes to apply to the entire stack.
	CustomJson string `json:"customJson,omitempty" yaml:"customJson,omitempty"`

	/*
	   Name of the availability zone where instances will be created by default.
	   Cannot be set when `vpc_id` is set.
	*/
	DefaultAvailabilityZone string `json:"defaultAvailabilityZone,omitempty" yaml:"defaultAvailabilityZone,omitempty"`

	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName string `json:"defaultSshKeyName,omitempty" yaml:"defaultSshKeyName,omitempty"`

	// Keyword representing the naming scheme that will be used for instance hostnames within this stack.
	HostnameTheme string `json:"hostnameTheme,omitempty" yaml:"hostnameTheme,omitempty"`
}
