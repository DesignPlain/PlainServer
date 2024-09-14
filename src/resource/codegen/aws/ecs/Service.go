package ecs

import types "libds/aws/types"

type Service struct {
	// Platform version on which to run your service. Only applicable for `launch_type` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`

	// Configuration block for deployment circuit breaker. See below.
	DeploymentCircuitBreaker types.Ecs_ServiceDeploymentCircuitBreaker `json:"deploymentCircuitBreaker,omitempty" yaml:"deploymentCircuitBreaker,omitempty"`

	// Lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	DeploymentMinimumHealthyPercent int `json:"deploymentMinimumHealthyPercent,omitempty" yaml:"deploymentMinimumHealthyPercent,omitempty"`

	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `force_new_deployment` is enabled. The maximum number of `ordered_placement_strategy` blocks is `5`. See below.
	OrderedPlacementStrategies []types.Ecs_ServiceOrderedPlacementStrategy `json:"orderedPlacementStrategies,omitempty" yaml:"orderedPlacementStrategies,omitempty"`

	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	IamRole string `json:"iamRole,omitempty" yaml:"iamRole,omitempty"`

	// Network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. See below.
	NetworkConfiguration types.Ecs_ServiceNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// Configuration block for deployment controller configuration. See below.
	DeploymentController types.Ecs_ServiceDeploymentController `json:"deploymentController,omitempty" yaml:"deploymentController,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger an in-place update (redeployment). Useful with `"plantimestamp()"`. When using the triggers property you also need to set the forceNewDeployment property to True.
	Triggers map[string]string `json:"triggers,omitempty" yaml:"triggers,omitempty"`

	// Number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
	DesiredCount int `json:"desiredCount,omitempty" yaml:"desiredCount,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether to enable Amazon ECS Exec for the tasks within the service.
	EnableExecuteCommand bool `json:"enableExecuteCommand,omitempty" yaml:"enableExecuteCommand,omitempty"`

	// Enable to delete a service even if it wasn't scaled down to zero tasks. It's only necessary to use this if the service uses the `REPLICA` scheduling strategy.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	HealthCheckGracePeriodSeconds int `json:"healthCheckGracePeriodSeconds,omitempty" yaml:"healthCheckGracePeriodSeconds,omitempty"`

	// Whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
	PropagateTags string `json:"propagateTags,omitempty" yaml:"propagateTags,omitempty"`

	// Configuration for a volume specified in the task definition as a volume that is configured at launch time. Currently, the only supported volume type is an Amazon EBS volume. See below.
	VolumeConfiguration types.Ecs_ServiceVolumeConfiguration `json:"volumeConfiguration,omitempty" yaml:"volumeConfiguration,omitempty"`

	// If `true`, this provider will wait for the service to reach a steady state (like [`aws ecs wait services-stable`](https://docs.aws.amazon.com/cli/latest/reference/ecs/wait/services-stable.html)) before continuing. Default `false`.
	WaitForSteadyState bool `json:"waitForSteadyState,omitempty" yaml:"waitForSteadyState,omitempty"`

	// Information about the CloudWatch alarms. See below.
	Alarms types.Ecs_ServiceAlarms `json:"alarms,omitempty" yaml:"alarms,omitempty"`

	// Capacity provider strategies to use for the service. Can be one or more. These can be updated without destroying and recreating the service only if `force_new_deployment = true` and not changing from 0 `capacity_provider_strategy` blocks to greater than 0, or vice versa. See below. Conflicts with `launch_type`.
	CapacityProviderStrategies []types.Ecs_ServiceCapacityProviderStrategy `json:"capacityProviderStrategies,omitempty" yaml:"capacityProviderStrategies,omitempty"`

	// ARN of an ECS cluster.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	/*
	   Name of the service (up to 255 letters, numbers, hyphens, and underscores)

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for load balancers. See below.
	LoadBalancers []types.Ecs_ServiceLoadBalancer `json:"loadBalancers,omitempty" yaml:"loadBalancers,omitempty"`

	// Scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [-Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy-](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
	SchedulingStrategy string `json:"schedulingStrategy,omitempty" yaml:"schedulingStrategy,omitempty"`

	// Family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
	TaskDefinition string `json:"taskDefinition,omitempty" yaml:"taskDefinition,omitempty"`

	/*
	   Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g., `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `ordered_placement_strategy` and `placement_constraints` updates.
	   When using the forceNewDeployment property you also need to configure the triggers property.
	*/
	ForceNewDeployment bool `json:"forceNewDeployment,omitempty" yaml:"forceNewDeployment,omitempty"`

	// Launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`. Conflicts with `capacity_provider_strategy`.
	LaunchType string `json:"launchType,omitempty" yaml:"launchType,omitempty"`

	// Rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `force_new_deployment` is enabled. Maximum number of `placement_constraints` is `10`. See below.
	PlacementConstraints []types.Ecs_ServicePlacementConstraint `json:"placementConstraints,omitempty" yaml:"placementConstraints,omitempty"`

	// ECS Service Connect configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace. See below.
	ServiceConnectConfiguration types.Ecs_ServiceServiceConnectConfiguration `json:"serviceConnectConfiguration,omitempty" yaml:"serviceConnectConfiguration,omitempty"`

	// Service discovery registries for the service. The maximum number of `service_registries` blocks is `1`. See below.
	ServiceRegistries types.Ecs_ServiceServiceRegistries `json:"serviceRegistries,omitempty" yaml:"serviceRegistries,omitempty"`

	// Upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
	DeploymentMaximumPercent int `json:"deploymentMaximumPercent,omitempty" yaml:"deploymentMaximumPercent,omitempty"`

	// Whether to enable Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags bool `json:"enableEcsManagedTags,omitempty" yaml:"enableEcsManagedTags,omitempty"`
}
