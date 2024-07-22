package types

type Imagebuilder_ImagePipelineSchedule struct {
	/*
	   Cron expression of how often the pipeline start condition is evaluated. For example, `cron(0 0 - - ? -)` is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as `cron(0 0 - - -)`, must be updated to the six field syntax. For more information, see the [Image Builder User Guide](https://docs.aws.amazon.com/imagebuilder/latest/userguide/cron-expressions.html).

	   The following arguments are optional:
	*/
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`

	// The timezone that applies to the scheduling expression. For example, "Etc/UTC", "America/Los_Angeles" in the [IANA timezone format](https://www.joda.org/joda-time/timezones.html). If not specified this defaults to UTC.
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`

	// Condition when the pipeline should trigger a new image build. Valid values are `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` and `EXPRESSION_MATCH_ONLY`. Defaults to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`.
	PipelineExecutionStartCondition string `json:"pipelineExecutionStartCondition,omitempty" yaml:"pipelineExecutionStartCondition,omitempty"`
}