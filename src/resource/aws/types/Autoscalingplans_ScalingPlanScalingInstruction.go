package types

type Autoscalingplans_ScalingPlanScalingInstruction struct {
	// Boolean controlling whether dynamic scaling by AWS Auto Scaling is disabled. Defaults to `false`.
	DisableDynamicScaling bool `json:"disableDynamicScaling,omitempty" yaml:"disableDynamicScaling,omitempty"`

	/*
	   Predefined load metric to use for predictive scaling. You must specify either `predefined_load_metric_specification` or `customized_load_metric_specification` when configuring predictive scaling.
	   More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_PredefinedLoadMetricSpecification.html).
	*/
	PredefinedLoadMetricSpecification Autoscalingplans_ScalingPlanScalingInstructionPredefinedLoadMetricSpecification `json:"predefinedLoadMetricSpecification,omitempty" yaml:"predefinedLoadMetricSpecification,omitempty"`

	/*
	   Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	   Valid values: `SetForecastCapacityToMaxCapacity`, `SetMaxCapacityAboveForecastCapacity`, `SetMaxCapacityToForecastCapacity`.
	*/
	PredictiveScalingMaxCapacityBehavior string `json:"predictiveScalingMaxCapacityBehavior,omitempty" yaml:"predictiveScalingMaxCapacityBehavior,omitempty"`

	// Size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	PredictiveScalingMaxCapacityBuffer int `json:"predictiveScalingMaxCapacityBuffer,omitempty" yaml:"predictiveScalingMaxCapacityBuffer,omitempty"`

	// Predictive scaling mode. Valid values: `ForecastAndScale`, `ForecastOnly`.
	PredictiveScalingMode string `json:"predictiveScalingMode,omitempty" yaml:"predictiveScalingMode,omitempty"`

	// Amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	ScheduledActionBufferTime int `json:"scheduledActionBufferTime,omitempty" yaml:"scheduledActionBufferTime,omitempty"`

	// Minimum capacity of the resource.
	MinCapacity int `json:"minCapacity,omitempty" yaml:"minCapacity,omitempty"`

	// ID of the resource. This string consists of the resource type and unique identifier.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Scalable dimension associated with the resource. Valid values: `autoscaling:autoScalingGroup:DesiredCapacity`, `dynamodb:index:ReadCapacityUnits`, `dynamodb:index:WriteCapacityUnits`, `dynamodb:table:ReadCapacityUnits`, `dynamodb:table:WriteCapacityUnits`, `ecs:service:DesiredCount`, `ec2:spot-fleet-request:TargetCapacity`, `rds:cluster:ReadReplicaCount`.
	ScalableDimension string `json:"scalableDimension,omitempty" yaml:"scalableDimension,omitempty"`

	// Controls whether a resource's externally created scaling policies are kept or replaced. Valid values: `KeepExternalPolicies`, `ReplaceExternalPolicies`. Defaults to `KeepExternalPolicies`.
	ScalingPolicyUpdateBehavior string `json:"scalingPolicyUpdateBehavior,omitempty" yaml:"scalingPolicyUpdateBehavior,omitempty"`

	/*
	   Structure that defines new target tracking configurations. Each of these structures includes a specific scaling metric and a target value for the metric, along with various parameters to use with dynamic scaling.
	   More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_TargetTrackingConfiguration.html).
	*/
	TargetTrackingConfigurations []Autoscalingplans_ScalingPlanScalingInstructionTargetTrackingConfiguration `json:"targetTrackingConfigurations,omitempty" yaml:"targetTrackingConfigurations,omitempty"`

	/*
	   Customized load metric to use for predictive scaling. You must specify either `customized_load_metric_specification` or `predefined_load_metric_specification` when configuring predictive scaling.
	   More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_CustomizedLoadMetricSpecification.html).
	*/
	CustomizedLoadMetricSpecification Autoscalingplans_ScalingPlanScalingInstructionCustomizedLoadMetricSpecification `json:"customizedLoadMetricSpecification,omitempty" yaml:"customizedLoadMetricSpecification,omitempty"`

	// Maximum capacity of the resource. The exception to this upper limit is if you specify a non-default setting for `predictive_scaling_max_capacity_behavior`.
	MaxCapacity int `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// Namespace of the AWS service. Valid values: `autoscaling`, `dynamodb`, `ecs`, `ec2`, `rds`.
	ServiceNamespace string `json:"serviceNamespace,omitempty" yaml:"serviceNamespace,omitempty"`
}
