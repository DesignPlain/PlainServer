package types

type Cloudbuild_getTriggerGithubPullRequest struct {
	// Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"]
	CommentControl string `json:"commentControl,omitempty" yaml:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`

	// Regex of branches to match.
	Branch string `json:"branch,omitempty" yaml:"branch,omitempty"`
}
