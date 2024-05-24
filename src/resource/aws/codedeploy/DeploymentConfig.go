package codedeploy

import types "DesignSphere_Server/src/resource/aws/types"

type DeploymentConfig struct {
	// A traffic_routing_config block. Traffic Routing Config is documented below.
	TrafficRoutingConfig types.Codedeploy_DeploymentConfigTrafficRoutingConfig `json:"trafficRoutingConfig,omitempty" yaml:"trafficRoutingConfig,omitempty"`

	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform string `json:"computePlatform,omitempty" yaml:"computePlatform,omitempty"`

	// The name of the deployment config.
	DeploymentConfigName string `json:"deploymentConfigName,omitempty" yaml:"deploymentConfigName,omitempty"`

	// A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts types.Codedeploy_DeploymentConfigMinimumHealthyHosts `json:"minimumHealthyHosts,omitempty" yaml:"minimumHealthyHosts,omitempty"`
}
