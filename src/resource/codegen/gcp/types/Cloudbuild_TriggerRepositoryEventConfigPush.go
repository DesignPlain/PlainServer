package types

type Cloudbuild_TriggerRepositoryEventConfigPush struct {
	// Regex of tags to match.  Specify only one of branch or tag.
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`

	// Regex of branches to match.  Specify only one of branch or tag.
	Branch string `json:"branch,omitempty" yaml:"branch,omitempty"`

	// When true, only trigger a build if the revision regex does NOT match the gitRef regex.
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`
}
