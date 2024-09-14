package appautoscaling

import types "libds/aws/types"

type ScheduledAction struct {
	// Identifier of the resource associated with the scheduled action. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScheduledAction.html)
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in `timezone`. Documentation can be found in the `Timezone` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScheduledAction.html)
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// Time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for `start_time` and `end_time`. Valid values are the [canonical names of the IANA time zones supported by Joda-Time](https://www.joda.org/joda-time/timezones.html), such as `Etc/GMT+9` or `Pacific/Tahiti`. Default is `UTC`.
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`

	// Namespace of the AWS service. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScheduledAction.html) Example: ecs
	ServiceNamespace string `json:"serviceNamespace,omitempty" yaml:"serviceNamespace,omitempty"`

	// Date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	// Name of the scheduled action.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Scalable dimension. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScheduledAction.html) Example: ecs:service:DesiredCount
	ScalableDimension string `json:"scalableDimension,omitempty" yaml:"scalableDimension,omitempty"`

	// New minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction types.Appautoscaling_ScheduledActionScalableTargetAction `json:"scalableTargetAction,omitempty" yaml:"scalableTargetAction,omitempty"`
}
