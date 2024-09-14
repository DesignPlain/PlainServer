package datasync

import types "libds/aws/types"

type Task struct {
	// Specifies a schedule used to periodically transfer files from a source to a destination location.
	Schedule types.Datasync_TaskSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Configuration block containing the configuration of a DataSync Task Report. See `task_report_config` below.
	TaskReportConfig types.Datasync_TaskTaskReportConfig `json:"taskReportConfig,omitempty" yaml:"taskReportConfig,omitempty"`

	// Amazon Resource Name (ARN) of destination DataSync Location.
	DestinationLocationArn string `json:"destinationLocationArn,omitempty" yaml:"destinationLocationArn,omitempty"`

	// Filter rules that determines which files to include in a task.
	Includes types.Datasync_TaskIncludes `json:"includes,omitempty" yaml:"includes,omitempty"`

	// Name of the DataSync Task.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amazon Resource Name (ARN) of source DataSync Location.
	SourceLocationArn string `json:"sourceLocationArn,omitempty" yaml:"sourceLocationArn,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Task. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Amazon Resource Name (ARN) of the CloudWatch Log Group that is used to monitor and log events in the sync task.
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn,omitempty" yaml:"cloudwatchLogGroupArn,omitempty"`

	// Filter rules that determines which files to exclude from a task.
	Excludes types.Datasync_TaskExcludes `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	// Configuration block containing option that controls the default behavior when you start an execution of this DataSync Task. For each individual task execution, you can override these options by specifying an overriding configuration in those executions.
	Options types.Datasync_TaskOptions `json:"options,omitempty" yaml:"options,omitempty"`
}
