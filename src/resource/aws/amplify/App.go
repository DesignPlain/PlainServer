package amplify

import types "DesignSphere_Server/src/resource/aws/types"

type App struct {
	// Credentials for basic authorization for an Amplify app.
	BasicAuthCredentials string `json:"basicAuthCredentials,omitempty" yaml:"basicAuthCredentials,omitempty"`

	// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
	BuildSpec string `json:"buildSpec,omitempty" yaml:"buildSpec,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Automated branch creation configuration for an Amplify app. An `auto_branch_creation_config` block is documented below.
	AutoBranchCreationConfig types.Amplify_AppAutoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty" yaml:"autoBranchCreationConfig,omitempty"`

	// The [custom HTTP headers](https://docs.aws.amazon.com/amplify/latest/userguide/custom-headers.html) for an Amplify app.
	CustomHeaders string `json:"customHeaders,omitempty" yaml:"customHeaders,omitempty"`

	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation bool `json:"enableAutoBranchCreation,omitempty" yaml:"enableAutoBranchCreation,omitempty"`

	// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	EnableBasicAuth bool `json:"enableBasicAuth,omitempty" yaml:"enableBasicAuth,omitempty"`

	// Name for an Amplify app.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	OauthToken string `json:"oauthToken,omitempty" yaml:"oauthToken,omitempty"`

	// Enables auto-building of branches for the Amplify App.
	EnableBranchAutoBuild bool `json:"enableBranchAutoBuild,omitempty" yaml:"enableBranchAutoBuild,omitempty"`

	// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion bool `json:"enableBranchAutoDeletion,omitempty" yaml:"enableBranchAutoDeletion,omitempty"`

	// Environment variables map for an Amplify app.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// Repository for an Amplify app.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// Automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns []string `json:"autoBranchCreationPatterns,omitempty" yaml:"autoBranchCreationPatterns,omitempty"`

	// Custom rewrite and redirect rules for an Amplify app. A `custom_rule` block is documented below.
	CustomRules []types.Amplify_AppCustomRule `json:"customRules,omitempty" yaml:"customRules,omitempty"`

	// Description for an Amplify app.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn string `json:"iamServiceRoleArn,omitempty" yaml:"iamServiceRoleArn,omitempty"`

	// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`
}
