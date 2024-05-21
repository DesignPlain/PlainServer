package firebaserules

type Release struct {
	// Format: `projects/{project_id}/releases/{release_id}`\Firestore Rules Releases will --always-- have the name 'cloud.firestore'
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.



	   - - -
	*/
	RulesetName string `json:"rulesetName,omitempty" yaml:"rulesetName,omitempty"`
}
