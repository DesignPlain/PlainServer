package types

type Datapipeline_PipelineDefinitionPipelineObject struct {
	// Configuration block for Key-value pairs that define the properties of the object. See below
	Fields []Datapipeline_PipelineDefinitionPipelineObjectField `json:"fields,omitempty" yaml:"fields,omitempty"`

	// ID of the object.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// ARN of the storage connector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
