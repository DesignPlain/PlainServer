package types

type Datapipeline_PipelineDefinitionParameterValue struct {
	// ID of the parameter value.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Field value, expressed as a String.
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`
}
