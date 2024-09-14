package types

type Codepipeline_CustomActionTypeOutputArtifactDetails struct {
	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	MaximumCount int `json:"maximumCount,omitempty" yaml:"maximumCount,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	MinimumCount int `json:"minimumCount,omitempty" yaml:"minimumCount,omitempty"`
}
