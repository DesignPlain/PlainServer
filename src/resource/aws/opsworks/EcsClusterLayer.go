package opsworks

import types "DesignSphere_Server/src/resource/aws/types"

type EcsClusterLayer struct {
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps bool `json:"autoAssignPublicIps,omitempty" yaml:"autoAssignPublicIps,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJson string `json:"customJson,omitempty" yaml:"customJson,omitempty"`

	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `json:"customSecurityGroupIds,omitempty" yaml:"customSecurityGroupIds,omitempty"`

	//
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" yaml:"customShutdownRecipes,omitempty"`

	//
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" yaml:"customUndeployRecipes,omitempty"`

	// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []types.Opsworks_EcsClusterLayerEbsVolume `json:"ebsVolumes,omitempty" yaml:"ebsVolumes,omitempty"`

	// The ECS Cluster ARN of the layer.
	EcsClusterArn string `json:"ecsClusterArn,omitempty" yaml:"ecsClusterArn,omitempty"`

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps bool `json:"autoAssignElasticIps,omitempty" yaml:"autoAssignElasticIps,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing bool `json:"autoHealing,omitempty" yaml:"autoHealing,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn string `json:"customInstanceProfileArn,omitempty" yaml:"customInstanceProfileArn,omitempty"`

	//
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" yaml:"customSetupRecipes,omitempty"`

	//
	LoadBasedAutoScaling types.Opsworks_EcsClusterLayerLoadBasedAutoScaling `json:"loadBasedAutoScaling,omitempty" yaml:"loadBasedAutoScaling,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances bool `json:"useEbsOptimizedInstances,omitempty" yaml:"useEbsOptimizedInstances,omitempty"`

	//
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" yaml:"customDeployRecipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown bool `json:"drainElbOnShutdown,omitempty" yaml:"drainElbOnShutdown,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" yaml:"installUpdatesOnBoot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout int `json:"instanceShutdownTimeout,omitempty" yaml:"instanceShutdownTimeout,omitempty"`

	// A human-readable name for the layer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `json:"systemPackages,omitempty" yaml:"systemPackages,omitempty"`

	/*
	   A mapping of tags to assign to the resource.

	   The following extra optional arguments, all lists of Chef recipe names, allow
	   custom Chef recipes to be applied to layer instances at the five different
	   lifecycle events, if custom cookbooks are enabled on the layer's stack:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	CloudwatchConfiguration types.Opsworks_EcsClusterLayerCloudwatchConfiguration `json:"cloudwatchConfiguration,omitempty" yaml:"cloudwatchConfiguration,omitempty"`

	//
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" yaml:"customConfigureRecipes,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer string `json:"elasticLoadBalancer,omitempty" yaml:"elasticLoadBalancer,omitempty"`

	// ID of the stack the layer will belong to.
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`
}
