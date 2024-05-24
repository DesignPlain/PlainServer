package cloudwatch

import types "DesignSphere_Server/src/resource/aws/types"

type CompositeAlarm struct {
	// The description for the composite alarm.
	AlarmDescription string `json:"alarmDescription,omitempty" yaml:"alarmDescription,omitempty"`

	// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
	AlarmRule string `json:"alarmRule,omitempty" yaml:"alarmRule,omitempty"`

	// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	InsufficientDataActions []string `json:"insufficientDataActions,omitempty" yaml:"insufficientDataActions,omitempty"`

	// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	OkActions []string `json:"okActions,omitempty" yaml:"okActions,omitempty"`

	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
	ActionsEnabled bool `json:"actionsEnabled,omitempty" yaml:"actionsEnabled,omitempty"`

	// Actions will be suppressed if the suppressor alarm is in the ALARM state.
	ActionsSuppressor types.Cloudwatch_CompositeAlarmActionsSuppressor `json:"actionsSuppressor,omitempty" yaml:"actionsSuppressor,omitempty"`

	// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	AlarmActions []string `json:"alarmActions,omitempty" yaml:"alarmActions,omitempty"`

	// The name for the composite alarm. This name must be unique within the region.
	AlarmName string `json:"alarmName,omitempty" yaml:"alarmName,omitempty"`
}
