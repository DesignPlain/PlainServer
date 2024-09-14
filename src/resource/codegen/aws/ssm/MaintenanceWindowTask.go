package ssm

import types "libds/aws/types"

type MaintenanceWindowTask struct {
	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" yaml:"serviceRoleArn,omitempty"`

	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []types.Ssm_MaintenanceWindowTaskTarget `json:"targets,omitempty" yaml:"targets,omitempty"`

	// Configuration block with parameters for task execution.
	TaskInvocationParameters types.Ssm_MaintenanceWindowTaskTaskInvocationParameters `json:"taskInvocationParameters,omitempty" yaml:"taskInvocationParameters,omitempty"`

	// The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
	TaskType string `json:"taskType,omitempty" yaml:"taskType,omitempty"`

	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency string `json:"maxConcurrency,omitempty" yaml:"maxConcurrency,omitempty"`

	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors string `json:"maxErrors,omitempty" yaml:"maxErrors,omitempty"`

	// The name of the maintenance window task.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The ARN of the task to execute.
	TaskArn string `json:"taskArn,omitempty" yaml:"taskArn,omitempty"`

	// The Id of the maintenance window to register the task with.
	WindowId string `json:"windowId,omitempty" yaml:"windowId,omitempty"`

	// Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are `CONTINUE_TASK` and `CANCEL_TASK`.
	CutoffBehavior string `json:"cutoffBehavior,omitempty" yaml:"cutoffBehavior,omitempty"`

	// The description of the maintenance window task.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
