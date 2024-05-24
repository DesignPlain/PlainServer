package ecs

import types "DesignSphere_Server/src/resource/aws/types"

type TaskSet struct {
	/*
	   The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.

	   The following arguments are optional:
	*/
	TaskDefinition string `json:"taskDefinition,omitempty" yaml:"taskDefinition,omitempty"`

	// Whether the provider should wait until the task set has reached `STEADY_STATE`.
	WaitUntilStable bool `json:"waitUntilStable,omitempty" yaml:"waitUntilStable,omitempty"`

	// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
	LaunchType string `json:"launchType,omitempty" yaml:"launchType,omitempty"`

	// A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The external ID associated with the task set.
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
	NetworkConfiguration types.Ecs_TaskSetNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// The platform version on which to run your service. Only applicable for `launch_type` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`

	// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
	Scale types.Ecs_TaskSetScale `json:"scale,omitempty" yaml:"scale,omitempty"`

	// The short name or ARN of the ECS service.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// The service discovery registries for the service. The maximum number of `service_registries` blocks is `1`. Detailed below.
	ServiceRegistries types.Ecs_TaskSetServiceRegistries `json:"serviceRegistries,omitempty" yaml:"serviceRegistries,omitempty"`

	// The short name or ARN of the cluster that hosts the service to create the task set in.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// Details on load balancers that are used with a task set. Detailed below.
	LoadBalancers []types.Ecs_TaskSetLoadBalancer `json:"loadBalancers,omitempty" yaml:"loadBalancers,omitempty"`

	// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
	WaitUntilStableTimeout string `json:"waitUntilStableTimeout,omitempty" yaml:"waitUntilStableTimeout,omitempty"`

	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies []types.Ecs_TaskSetCapacityProviderStrategy `json:"capacityProviderStrategies,omitempty" yaml:"capacityProviderStrategies,omitempty"`

	// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`
}
