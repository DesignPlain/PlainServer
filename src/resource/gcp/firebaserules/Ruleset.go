package firebaserules

import types "DesignSphere_Server/src/resource/gcp/types"

type Ruleset struct {
	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// `Source` for the `Ruleset`.
	Source types.Firebaserules_RulesetSource `json:"source,omitempty" yaml:"source,omitempty"`
}
