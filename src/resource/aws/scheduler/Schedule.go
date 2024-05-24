package scheduler

import types "DesignSphere_Server/src/resource/aws/types"

type Schedule struct {
	// Name of the schedule group to associate with this schedule. When omitted, the `default` schedule group is used.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// ARN for the customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Defines when the schedule runs. Read more in [Schedule types on EventBridge Scheduler](https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html).
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`

	// The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the start date you specify. EventBridge Scheduler ignores the start date for one-time schedules. Example: `2030-01-01T01:00:00Z`.
	StartDate string `json:"startDate,omitempty" yaml:"startDate,omitempty"`

	// Specifies whether the schedule is enabled or disabled. One of: `ENABLED` (default), `DISABLED`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   Configures the target of the schedule. Detailed below.

	   The following arguments are optional:
	*/
	Target types.Scheduler_ScheduleTarget `json:"target,omitempty" yaml:"target,omitempty"`

	// The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the end date you specify. EventBridge Scheduler ignores the end date for one-time schedules. Example: `2030-01-01T01:00:00Z`.
	EndDate string `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	// Configures a time window during which EventBridge Scheduler invokes the schedule. Detailed below.
	FlexibleTimeWindow types.Scheduler_ScheduleFlexibleTimeWindow `json:"flexibleTimeWindow,omitempty" yaml:"flexibleTimeWindow,omitempty"`

	// Name of the schedule. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Timezone in which the scheduling expression is evaluated. Defaults to `UTC`. Example: `Australia/Sydney`.
	ScheduleExpressionTimezone string `json:"scheduleExpressionTimezone,omitempty" yaml:"scheduleExpressionTimezone,omitempty"`

	// Brief description of the schedule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
