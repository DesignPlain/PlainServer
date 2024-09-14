package autoscaling

import types "libds/aws/types"

type Group struct {
	// Allows deleting the Auto Scaling Group without waiting for all instances in the warm pool to terminate.
	ForceDeleteWarmPool bool `json:"forceDeleteWarmPool,omitempty" yaml:"forceDeleteWarmPool,omitempty"`

	// Configuration block containing settings to define launch targets for Auto Scaling groups. See Mixed Instances Policy below for more details.
	MixedInstancesPolicy types.Autoscaling_GroupMixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty" yaml:"mixedInstancesPolicy,omitempty"`

	/*
	   If this block is configured, add a [Warm Pool](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html)
	   to the specified Auto Scaling group. Defined below
	*/
	WarmPool types.Autoscaling_GroupWarmPool `json:"warmPool,omitempty" yaml:"warmPool,omitempty"`

	// Reserved.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`

	// Whether to ignore failed [Auto Scaling scaling activities](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-verify-scaling-activity.html) while waiting for capacity. The default is `false` -- failed scaling activities cause errors to be returned.
	IgnoreFailedScalingActivities bool `json:"ignoreFailedScalingActivities,omitempty" yaml:"ignoreFailedScalingActivities,omitempty"`

	/*
	   If this block is configured, start an
	   [Instance Refresh](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html)
	   when this Auto Scaling Group is updated. Defined below.
	*/
	InstanceRefresh types.Autoscaling_GroupInstanceRefresh `json:"instanceRefresh,omitempty" yaml:"instanceRefresh,omitempty"`

	/*
	   List of elastic load balancer names to add to the autoscaling
	   group names. Only valid for classic load balancers. For ALBs, use `target_group_arns` instead. To remove all load balancer attachments an empty list should be specified.
	*/
	LoadBalancers []string `json:"loadBalancers,omitempty" yaml:"loadBalancers,omitempty"`

	// Name of the placement group into which you'll launch your instances, if any.
	PlacementGroup string `json:"placementGroup,omitempty" yaml:"placementGroup,omitempty"`

	// ARN of the service-linked role that the ASG will use to call other AWS services
	ServiceLinkedRoleArn string `json:"serviceLinkedRoleArn,omitempty" yaml:"serviceLinkedRoleArn,omitempty"`

	// List of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with `availability_zones`.
	VpcZoneIdentifiers []string `json:"vpcZoneIdentifiers,omitempty" yaml:"vpcZoneIdentifiers,omitempty"`

	// A list of Availability Zones where instances in the Auto Scaling group can be created. Used for launching into the default VPC subnet in each Availability Zone when not using the `vpc_zone_identifier` attribute, or for attaching a network interface when an existing network interface ID is specified in a launch template. Conflicts with `vpc_zone_identifier`.
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// Amount of time, in seconds, until a newly launched instance can contribute to the Amazon CloudWatch metrics. This delay lets an instance finish initializing before Amazon EC2 Auto Scaling aggregates instance metrics, resulting in more reliable usage data. Set this value equal to the amount of time that it takes for resource consumption to become stable after an instance reaches the InService state. (See [Set the default instance warmup for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html))
	DefaultInstanceWarmup int `json:"defaultInstanceWarmup,omitempty" yaml:"defaultInstanceWarmup,omitempty"`

	// The unit of measurement for the value specified for `desired_capacity`. Supported for attribute-based instance type selection only. Valid values: `"units"`, `"vcpu"`, `"memory-mib"`.
	DesiredCapacityType string `json:"desiredCapacityType,omitempty" yaml:"desiredCapacityType,omitempty"`

	/*
	   Creates a unique name beginning with the specified
	   prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   Whether newly launched instances
	   are automatically protected from termination by Amazon EC2 Auto Scaling when
	   scaling in. For more information about preventing instances from terminating
	   on scale in, see [Using instance scale-in protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html)
	   in the Amazon EC2 Auto Scaling User Guide.
	*/
	ProtectFromScaleIn bool `json:"protectFromScaleIn,omitempty" yaml:"protectFromScaleIn,omitempty"`

	// List of policies to decide how the instances in the Auto Scaling Group should be terminated. The allowed values are `OldestInstance`, `NewestInstance`, `OldestLaunchConfiguration`, `ClosestToNextInstanceHour`, `OldestLaunchTemplate`, `AllocationStrategy`, `Default`. Additionally, the ARN of a Lambda function can be specified for custom termination policies.
	TerminationPolicies []string `json:"terminationPolicies,omitempty" yaml:"terminationPolicies,omitempty"`

	/*
	   Maximum
	   [duration](https://golang.org/pkg/time/#ParseDuration) that the provider should
	   wait for ASG instances to be healthy before timing out. (See also Waiting
	   for Capacity below.) Setting this to "0" causes
	   the provider to skip all Capacity Waiting behavior.
	*/
	WaitForCapacityTimeout string `json:"waitForCapacityTimeout,omitempty" yaml:"waitForCapacityTimeout,omitempty"`

	/*
	   Setting this will cause Pulumi to wait
	   for exactly this number of healthy instances from this Auto Scaling Group in
	   all attached load balancers on both create and update operations. (Takes
	   precedence over `min_elb_capacity` behavior.)
	   (See also Waiting for Capacity below.)
	*/
	WaitForElbCapacity int `json:"waitForElbCapacity,omitempty" yaml:"waitForElbCapacity,omitempty"`

	// List of metrics to collect. The allowed values are defined by the [underlying AWS API](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html).
	EnabledMetrics []string `json:"enabledMetrics,omitempty" yaml:"enabledMetrics,omitempty"`

	// Maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 86400 and 31536000 seconds.
	MaxInstanceLifetime int `json:"maxInstanceLifetime,omitempty" yaml:"maxInstanceLifetime,omitempty"`

	/*
	   Setting this causes Pulumi to wait for
	   this number of instances from this Auto Scaling Group to show up healthy in the
	   ELB only on creation. Updates will not wait on ELB instance number changes.
	   (See also Waiting for Capacity below.)
	*/
	MinElbCapacity int `json:"minElbCapacity,omitempty" yaml:"minElbCapacity,omitempty"`

	// Granularity to associate with the metrics to collect. The only valid value is `1Minute`. Default is `1Minute`.
	MetricsGranularity string `json:"metricsGranularity,omitempty" yaml:"metricsGranularity,omitempty"`

	/*
	   List of processes to suspend for the Auto Scaling Group. The allowed values are `Launch`, `Terminate`, `HealthCheck`, `ReplaceUnhealthy`, `AZRebalance`, `AlarmNotification`, `ScheduledActions`, `AddToLoadBalancer`, `InstanceRefresh`.
	   Note that if you suspend either the `Launch` or `Terminate` process types, it can prevent your Auto Scaling Group from functioning properly.
	*/
	SuspendedProcesses []string `json:"suspendedProcesses,omitempty" yaml:"suspendedProcesses,omitempty"`

	// Configuration block(s) containing resource tags. See Tag below for more details.
	Tags []types.Autoscaling_GroupTag `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Minimum size of the Auto Scaling Group.
	   (See also Waiting for Capacity below.)
	*/
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// Set of `aws.alb.TargetGroup` ARNs, for use with Application or Network Load Balancing. To remove all target group attachments an empty list should be specified.
	TargetGroupArns []string `json:"targetGroupArns,omitempty" yaml:"targetGroupArns,omitempty"`

	// Attaches one or more traffic sources to the specified Auto Scaling group.
	TrafficSources []types.Autoscaling_GroupTrafficSource `json:"trafficSources,omitempty" yaml:"trafficSources,omitempty"`

	/*
	   Allows deleting the Auto Scaling Group without waiting
	   for all instances in the pool to terminate. You can force an Auto Scaling Group to delete
	   even if it's in the process of scaling a resource. Normally, this provider
	   drains all the instances before deleting the group. This bypasses that
	   behavior and potentially leaves resources dangling.
	*/
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// Time (in seconds) after instance comes into service before checking health.
	HealthCheckGracePeriod int `json:"healthCheckGracePeriod,omitempty" yaml:"healthCheckGracePeriod,omitempty"`

	// If this block is configured, add a instance maintenance policy to the specified Auto Scaling group. Defined below.
	InstanceMaintenancePolicy types.Autoscaling_GroupInstanceMaintenancePolicy `json:"instanceMaintenancePolicy,omitempty" yaml:"instanceMaintenancePolicy,omitempty"`

	/*
	   One or more
	   [Lifecycle Hooks](http://docs.aws.amazon.com/autoscaling/latest/userguide/lifecycle-hooks.html)
	   to attach to the Auto Scaling Group --before-- instances are launched. The
	   syntax is exactly the same as the separate
	   `aws.autoscaling.LifecycleHook`
	   resource, without the `autoscaling_group_name` attribute. Please note that this will only work when creating
	   a new Auto Scaling Group. For all other use-cases, please use `aws.autoscaling.LifecycleHook` resource.
	*/
	InitialLifecycleHooks []types.Autoscaling_GroupInitialLifecycleHook `json:"initialLifecycleHooks,omitempty" yaml:"initialLifecycleHooks,omitempty"`

	// Name of the launch configuration to use.
	LaunchConfiguration string `json:"launchConfiguration,omitempty" yaml:"launchConfiguration,omitempty"`

	// Maximum size of the Auto Scaling Group.
	MaxSize int `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`

	// Name of the Auto Scaling Group. By default generated by Pulumi. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	DefaultCooldown int `json:"defaultCooldown,omitempty" yaml:"defaultCooldown,omitempty"`

	/*
	   Number of Amazon EC2 instances that
	   should be running in the group. (See also Waiting for
	   Capacity below.)
	*/
	DesiredCapacity int `json:"desiredCapacity,omitempty" yaml:"desiredCapacity,omitempty"`

	// "EC2" or "ELB". Controls how health checking is done.
	HealthCheckType string `json:"healthCheckType,omitempty" yaml:"healthCheckType,omitempty"`

	// Whether capacity rebalance is enabled. Otherwise, capacity rebalance is disabled.
	CapacityRebalance bool `json:"capacityRebalance,omitempty" yaml:"capacityRebalance,omitempty"`

	// Nested argument with Launch template specification to use to launch instances. See Launch Template below for more details.
	LaunchTemplate types.Autoscaling_GroupLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`
}
