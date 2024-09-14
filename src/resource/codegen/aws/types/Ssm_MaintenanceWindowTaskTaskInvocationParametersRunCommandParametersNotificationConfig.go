package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig struct {
	// An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
	NotificationArn string `json:"notificationArn,omitempty" yaml:"notificationArn,omitempty"`

	// The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
	NotificationEvents []string `json:"notificationEvents,omitempty" yaml:"notificationEvents,omitempty"`

	// When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`
	NotificationType string `json:"notificationType,omitempty" yaml:"notificationType,omitempty"`
}
