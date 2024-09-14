package types

type Amplify_AppAutoBranchCreationConfig struct {
	// Enables pull request previews for the autocreated branch.
	EnablePullRequestPreview bool `json:"enablePullRequestPreview,omitempty" yaml:"enablePullRequestPreview,omitempty"`

	// Environment variables for the autocreated branch.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// Framework for the autocreated branch.
	Framework string `json:"framework,omitempty" yaml:"framework,omitempty"`

	// Basic authorization credentials for the autocreated branch.
	BasicAuthCredentials string `json:"basicAuthCredentials,omitempty" yaml:"basicAuthCredentials,omitempty"`

	// Enables basic authorization for the autocreated branch.
	EnableBasicAuth bool `json:"enableBasicAuth,omitempty" yaml:"enableBasicAuth,omitempty"`

	// Enables performance mode for the branch.
	EnablePerformanceMode bool `json:"enablePerformanceMode,omitempty" yaml:"enablePerformanceMode,omitempty"`

	// Amplify environment name for the pull request.
	PullRequestEnvironmentName string `json:"pullRequestEnvironmentName,omitempty" yaml:"pullRequestEnvironmentName,omitempty"`

	// Describes the current stage for the autocreated branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`

	// Build specification (build spec) for the autocreated branch.
	BuildSpec string `json:"buildSpec,omitempty" yaml:"buildSpec,omitempty"`

	// Enables auto building for the autocreated branch.
	EnableAutoBuild bool `json:"enableAutoBuild,omitempty" yaml:"enableAutoBuild,omitempty"`
}
