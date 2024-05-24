package glue

import types "DesignSphere_Server/src/resource/aws/types"

type Trigger struct {
	// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A description of the new trigger.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Start the trigger. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingConditions []types.Glue_TriggerEventBatchingCondition `json:"eventBatchingConditions,omitempty" yaml:"eventBatchingConditions,omitempty"`

	// The name of the trigger.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
	StartOnCreation bool `json:"startOnCreation,omitempty" yaml:"startOnCreation,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []types.Glue_TriggerAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate types.Glue_TriggerPredicate `json:"predicate,omitempty" yaml:"predicate,omitempty"`

	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName string `json:"workflowName,omitempty" yaml:"workflowName,omitempty"`
}
