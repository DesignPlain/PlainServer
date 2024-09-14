package opsworks

import types "libds/aws/types"

type CustomLayer struct {
	//
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" yaml:"customSetupRecipes,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See EBS Volume.
	EbsVolumes []types.Opsworks_CustomLayerEbsVolume `json:"ebsVolumes,omitempty" yaml:"ebsVolumes,omitempty"`

	// A human-readable name for the layer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `json:"systemPackages,omitempty" yaml:"systemPackages,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps bool `json:"autoAssignPublicIps,omitempty" yaml:"autoAssignPublicIps,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See Cloudwatch Configuration.
	CloudwatchConfiguration types.Opsworks_CustomLayerCloudwatchConfiguration `json:"cloudwatchConfiguration,omitempty" yaml:"cloudwatchConfiguration,omitempty"`

	//
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" yaml:"customConfigureRecipes,omitempty"`

	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances bool `json:"useEbsOptimizedInstances,omitempty" yaml:"useEbsOptimizedInstances,omitempty"`

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps bool `json:"autoAssignElasticIps,omitempty" yaml:"autoAssignElasticIps,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJson string `json:"customJson,omitempty" yaml:"customJson,omitempty"`

	//
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" yaml:"customUndeployRecipes,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout int `json:"instanceShutdownTimeout,omitempty" yaml:"instanceShutdownTimeout,omitempty"`

	// Load-based auto scaling configuration. See Load Based AutoScaling
	LoadBasedAutoScaling types.Opsworks_CustomLayerLoadBasedAutoScaling `json:"loadBasedAutoScaling,omitempty" yaml:"loadBasedAutoScaling,omitempty"`

	// ID of the stack the layer will belong to.
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`

	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   The following extra optional arguments, all lists of Chef recipe names, allow
	   custom Chef recipes to be applied to layer instances at the five different
	   lifecycle events, if custom cookbooks are enabled on the layer's stack:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing bool `json:"autoHealing,omitempty" yaml:"autoHealing,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn string `json:"customInstanceProfileArn,omitempty" yaml:"customInstanceProfileArn,omitempty"`

	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `json:"customSecurityGroupIds,omitempty" yaml:"customSecurityGroupIds,omitempty"`

	//
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" yaml:"customShutdownRecipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown bool `json:"drainElbOnShutdown,omitempty" yaml:"drainElbOnShutdown,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer string `json:"elasticLoadBalancer,omitempty" yaml:"elasticLoadBalancer,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" yaml:"installUpdatesOnBoot,omitempty"`

	//
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" yaml:"customDeployRecipes,omitempty"`
}
