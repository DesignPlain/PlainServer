package types

type Datapipeline_getPipelineDefinitionPipelineObjectField struct {
	// Field identifier.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Field value, expressed as the identifier of another object
	RefValue string `json:"refValue,omitempty" yaml:"refValue,omitempty"`

	// Field value, expressed as a String.
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`
}
