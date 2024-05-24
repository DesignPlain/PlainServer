package types

type Codebuild_WebhookFilterGroupFilter struct {
	// For a filter that uses `EVENT` type, a comma-separated string that specifies one event: `PUSH`, `PULL_REQUEST_CREATED`, `PULL_REQUEST_UPDATED`, `PULL_REQUEST_REOPENED`. `PULL_REQUEST_MERGED` works with GitHub & GitHub Enterprise only. For a filter that uses any of the other filter types, a regular expression.
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// The webhook filter group's type. Valid values for this parameter are: `EVENT`, `BASE_REF`, `HEAD_REF`, `ACTOR_ACCOUNT_ID`, `FILE_PATH`, `COMMIT_MESSAGE`. At least one filter group must specify `EVENT` as its type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// If set to `true`, the specified filter does -not- trigger a build. Defaults to `false`.
	ExcludeMatchedPattern bool `json:"excludeMatchedPattern,omitempty" yaml:"excludeMatchedPattern,omitempty"`
}
