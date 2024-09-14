package types

type Codepipeline_PipelineVariable struct {
	// The default value of a pipeline-level variable.
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	/*
	   The description of a pipeline-level variable.

	   > --Note:-- The input artifact of an action must exactly match the output artifact declared in a preceding action, but the input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of a pipeline-level variable.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
