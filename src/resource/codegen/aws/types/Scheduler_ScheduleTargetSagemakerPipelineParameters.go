package types

type Scheduler_ScheduleTargetSagemakerPipelineParameters struct {
	// Set of up to 200 parameter names and values to use when executing the SageMaker Model Building Pipeline. Detailed below.
	PipelineParameters []Scheduler_ScheduleTargetSagemakerPipelineParametersPipelineParameter `json:"pipelineParameters,omitempty" yaml:"pipelineParameters,omitempty"`
}
