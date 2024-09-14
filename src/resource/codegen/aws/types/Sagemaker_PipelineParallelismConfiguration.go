package types

type Sagemaker_PipelineParallelismConfiguration struct {
	// The max number of steps that can be executed in parallel.
	MaxParallelExecutionSteps int `json:"maxParallelExecutionSteps,omitempty" yaml:"maxParallelExecutionSteps,omitempty"`
}
