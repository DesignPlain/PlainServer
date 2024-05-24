package types

type Datapipeline_PipelineDefinitionParameterObject struct {
	// Configuration block for attributes of the parameter object. See below
	Attributes []Datapipeline_PipelineDefinitionParameterObjectAttribute `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// ID of the parameter object.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
