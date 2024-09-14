package types

type Codebuild_WebhookScopeConfiguration struct {
	// The domain of the GitHub Enterprise organization. Required if your project's source type is GITHUB_ENTERPRISE.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// The name of either the enterprise or organization.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of scope for a GitHub webhook. Valid values for this parameter are: `GITHUB_ORGANIZATION`, `GITHUB_GLOBAL`.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
