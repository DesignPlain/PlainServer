package redshift

import types "DesignSphere_Server/src/resource/aws/types"

type ScheduledAction struct {
	// Whether to enable the scheduled action. Default is `true` .
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	// The IAM role to assume to run the scheduled action.
	IamRole string `json:"iamRole,omitempty" yaml:"iamRole,omitempty"`

	// The scheduled action name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? - MON -)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// Target action. Documented below.
	TargetAction types.Redshift_ScheduledActionTargetAction `json:"targetAction,omitempty" yaml:"targetAction,omitempty"`

	// The description of the scheduled action.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
