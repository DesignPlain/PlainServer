package amplify

type Branch struct {
	// Basic authorization credentials for the branch.
	BasicAuthCredentials string `json:"basicAuthCredentials,omitempty" yaml:"basicAuthCredentials,omitempty"`

	// Name for the branch.
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// Description for the branch.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Enables pull request previews for this branch.
	EnablePullRequestPreview bool `json:"enablePullRequestPreview,omitempty" yaml:"enablePullRequestPreview,omitempty"`

	// Framework for the branch.
	Framework string `json:"framework,omitempty" yaml:"framework,omitempty"`

	// Unique ID for an Amplify app.
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	// Enables auto building for the branch.
	EnableAutoBuild bool `json:"enableAutoBuild,omitempty" yaml:"enableAutoBuild,omitempty"`

	// Enables performance mode for the branch.
	EnablePerformanceMode bool `json:"enablePerformanceMode,omitempty" yaml:"enablePerformanceMode,omitempty"`

	// ARN for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn string `json:"backendEnvironmentArn,omitempty" yaml:"backendEnvironmentArn,omitempty"`

	// Enables basic authorization for the branch.
	EnableBasicAuth bool `json:"enableBasicAuth,omitempty" yaml:"enableBasicAuth,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Content Time To Live (TTL) for the website in seconds.
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	// Display name for a branch. This is used as the default domain prefix.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Enables notifications for the branch.
	EnableNotification bool `json:"enableNotification,omitempty" yaml:"enableNotification,omitempty"`

	// Environment variables for the branch.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// Amplify environment name for the pull request.
	PullRequestEnvironmentName string `json:"pullRequestEnvironmentName,omitempty" yaml:"pullRequestEnvironmentName,omitempty"`

	// Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`
}
