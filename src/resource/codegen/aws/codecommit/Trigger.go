package codecommit

import types "libds/aws/types"

type Trigger struct {
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`

	// The name of the trigger.
	Triggers []types.Codecommit_TriggerTrigger `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}
