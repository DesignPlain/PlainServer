package opsworks

import types "DesignSphere_Server/src/resource/aws/types"

type RailsAppLayer struct {
	// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []types.Opsworks_RailsAppLayerEbsVolume `json:"ebsVolumes,omitempty" yaml:"ebsVolumes,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout int `json:"instanceShutdownTimeout,omitempty" yaml:"instanceShutdownTimeout,omitempty"`

	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   The following extra optional arguments, all lists of Chef recipe names, allow
	   custom Chef recipes to be applied to layer instances at the five different
	   lifecycle events, if custom cookbooks are enabled on the layer's stack:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances bool `json:"useEbsOptimizedInstances,omitempty" yaml:"useEbsOptimizedInstances,omitempty"`

	// When OpsWorks is managing Bundler, which version to use. Defaults to "1.5.3".
	BundlerVersion string `json:"bundlerVersion,omitempty" yaml:"bundlerVersion,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `json:"systemPackages,omitempty" yaml:"systemPackages,omitempty"`

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps bool `json:"autoAssignElasticIps,omitempty" yaml:"autoAssignElasticIps,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing bool `json:"autoHealing,omitempty" yaml:"autoHealing,omitempty"`

	//
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" yaml:"customConfigureRecipes,omitempty"`

	//
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" yaml:"customShutdownRecipes,omitempty"`

	// The version of Passenger to use. Defaults to "4.0.46".
	PassengerVersion string `json:"passengerVersion,omitempty" yaml:"passengerVersion,omitempty"`

	// The version of RubyGems to use. Defaults to "2.2.2".
	RubygemsVersion string `json:"rubygemsVersion,omitempty" yaml:"rubygemsVersion,omitempty"`

	// Keyword for the app server to use. Defaults to "apache_passenger".
	AppServer string `json:"appServer,omitempty" yaml:"appServer,omitempty"`

	//
	CloudwatchConfiguration types.Opsworks_RailsAppLayerCloudwatchConfiguration `json:"cloudwatchConfiguration,omitempty" yaml:"cloudwatchConfiguration,omitempty"`

	//
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" yaml:"customDeployRecipes,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJson string `json:"customJson,omitempty" yaml:"customJson,omitempty"`

	//
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" yaml:"customSetupRecipes,omitempty"`

	// The version of Ruby to use. Defaults to "2.0.0".
	RubyVersion string `json:"rubyVersion,omitempty" yaml:"rubyVersion,omitempty"`

	// ID of the stack the layer will belong to.
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`

	//
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" yaml:"customUndeployRecipes,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" yaml:"installUpdatesOnBoot,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer string `json:"elasticLoadBalancer,omitempty" yaml:"elasticLoadBalancer,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps bool `json:"autoAssignPublicIps,omitempty" yaml:"autoAssignPublicIps,omitempty"`

	//
	LoadBasedAutoScaling types.Opsworks_RailsAppLayerLoadBasedAutoScaling `json:"loadBasedAutoScaling,omitempty" yaml:"loadBasedAutoScaling,omitempty"`

	// Whether OpsWorks should manage bundler. On by default.
	ManageBundler bool `json:"manageBundler,omitempty" yaml:"manageBundler,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn string `json:"customInstanceProfileArn,omitempty" yaml:"customInstanceProfileArn,omitempty"`

	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `json:"customSecurityGroupIds,omitempty" yaml:"customSecurityGroupIds,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown bool `json:"drainElbOnShutdown,omitempty" yaml:"drainElbOnShutdown,omitempty"`

	// A human-readable name for the layer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
