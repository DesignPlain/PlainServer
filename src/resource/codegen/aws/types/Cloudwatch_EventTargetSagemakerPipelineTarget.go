package types

type Cloudwatch_EventTargetSagemakerPipelineTarget struct {
	// List of Parameter names and values for SageMaker Model Building Pipeline execution.
	PipelineParameterLists []Cloudwatch_EventTargetSagemakerPipelineTargetPipelineParameterList `json:"pipelineParameterLists,omitempty" yaml:"pipelineParameterLists,omitempty"`
}
