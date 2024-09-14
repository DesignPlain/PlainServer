package types

type Cloudbuild_getTriggerRepositoryEventConfigPush struct {
	/*
	   Regex of branches to match.

	   The syntax of the regular expressions accepted is the syntax accepted by
	   RE2 and described at https://github.com/google/re2/wiki/Syntax
	*/
	Branch string `json:"branch,omitempty" yaml:"branch,omitempty"`

	// If true, only trigger a build if the revision regex does NOT match the git_ref regex.
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`

	/*
	   Regex of tags to match.

	   The syntax of the regular expressions accepted is the syntax accepted by
	   RE2 and described at https://github.com/google/re2/wiki/Syntax
	*/
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`
}
