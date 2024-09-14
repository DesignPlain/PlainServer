package codedeploy

import types "libds/aws/types"

type DeploymentConfig struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform string `json:"computePlatform,omitempty" yaml:"computePlatform,omitempty"`

	// The name of the deployment config.
	DeploymentConfigName string `json:"deploymentConfigName,omitempty" yaml:"deploymentConfigName,omitempty"`

	// A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts types.Codedeploy_DeploymentConfigMinimumHealthyHosts `json:"minimumHealthyHosts,omitempty" yaml:"minimumHealthyHosts,omitempty"`

	// A traffic_routing_config block. Traffic Routing Config is documented below.
	TrafficRoutingConfig types.Codedeploy_DeploymentConfigTrafficRoutingConfig `json:"trafficRoutingConfig,omitempty" yaml:"trafficRoutingConfig,omitempty"`
}
