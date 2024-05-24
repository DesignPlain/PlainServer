package autoscaling

type LifecycleHook struct {
	// ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn string `json:"notificationTargetArn,omitempty" yaml:"notificationTargetArn,omitempty"`

	// ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty" yaml:"autoscalingGroupName,omitempty"`

	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult string `json:"defaultResult,omitempty" yaml:"defaultResult,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout int `json:"heartbeatTimeout,omitempty" yaml:"heartbeatTimeout,omitempty"`

	// Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
	LifecycleTransition string `json:"lifecycleTransition,omitempty" yaml:"lifecycleTransition,omitempty"`

	// Name of the lifecycle hook.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata string `json:"notificationMetadata,omitempty" yaml:"notificationMetadata,omitempty"`
}
