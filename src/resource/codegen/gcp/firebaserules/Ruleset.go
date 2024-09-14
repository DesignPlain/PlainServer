package firebaserules

import types "libds/gcp/types"

type Ruleset struct {
	// `Source` for the `Ruleset`.
	Source types.Firebaserules_RulesetSource `json:"source,omitempty" yaml:"source,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
