package types

type Cloudwatch_EventTargetSagemakerPipelineTargetPipelineParameterList struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
