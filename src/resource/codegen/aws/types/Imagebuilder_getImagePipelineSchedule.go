package types

type Imagebuilder_getImagePipelineSchedule struct {
	// Condition when the pipeline should trigger a new image build.
	PipelineExecutionStartCondition string `json:"pipelineExecutionStartCondition,omitempty" yaml:"pipelineExecutionStartCondition,omitempty"`

	// Cron expression of how often the pipeline start condition is evaluated.
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`
}
