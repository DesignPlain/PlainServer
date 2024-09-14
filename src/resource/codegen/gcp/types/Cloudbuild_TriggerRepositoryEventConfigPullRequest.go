package types

type Cloudbuild_TriggerRepositoryEventConfigPullRequest struct {
	/*
	   Regex of branches to match.
	   The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
	*/
	Branch string `json:"branch,omitempty" yaml:"branch,omitempty"`

	/*
	   Configure builds to run whether a repository owner or collaborator need to comment /gcbrun.
	   Possible values are: `COMMENTS_DISABLED`, `COMMENTS_ENABLED`, `COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY`.
	*/
	CommentControl string `json:"commentControl,omitempty" yaml:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`
}
