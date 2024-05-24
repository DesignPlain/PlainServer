package appautoscaling

import types "DesignSphere_Server/src/resource/aws/types"

type Policy struct {
	// AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
	ServiceNamespace string `json:"serviceNamespace,omitempty" yaml:"serviceNamespace,omitempty"`

	// Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration types.Appautoscaling_PolicyStepScalingPolicyConfiguration `json:"stepScalingPolicyConfiguration,omitempty" yaml:"stepScalingPolicyConfiguration,omitempty"`

	// Target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration types.Appautoscaling_PolicyTargetTrackingScalingPolicyConfiguration `json:"targetTrackingScalingPolicyConfiguration,omitempty" yaml:"targetTrackingScalingPolicyConfiguration,omitempty"`

	// Name of the policy. Must be between 1 and 255 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
	PolicyType string `json:"policyType,omitempty" yaml:"policyType,omitempty"`

	// Resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
	ScalableDimension string `json:"scalableDimension,omitempty" yaml:"scalableDimension,omitempty"`
}
