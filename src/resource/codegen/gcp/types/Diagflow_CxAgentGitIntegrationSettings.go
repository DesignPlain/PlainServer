package types

type Diagflow_CxAgentGitIntegrationSettings struct {
	/*
	   Settings of integration with GitHub.
	   Structure is documented below.
	*/
	GithubSettings Diagflow_CxAgentGitIntegrationSettingsGithubSettings `json:"githubSettings,omitempty" yaml:"githubSettings,omitempty"`
}
