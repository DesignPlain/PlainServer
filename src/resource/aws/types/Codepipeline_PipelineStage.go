package types

type Codepipeline_PipelineStage struct {
	// The name of the stage.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The action(s) to include in the stage. Defined as an `action` block below
	Actions []Codepipeline_PipelineStageAction `json:"actions,omitempty" yaml:"actions,omitempty"`
}
