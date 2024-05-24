package types

type Datapipeline_getPipelineDefinitionPipelineObject struct {
	// ARN of the storage connector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value pairs that define the properties of the object. See below
	Fields []Datapipeline_getPipelineDefinitionPipelineObjectField `json:"fields,omitempty" yaml:"fields,omitempty"`

	// ID of the object.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
