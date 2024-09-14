package types

type Codebuild_ProjectSource struct {
	// Configuration block. Detailed below.
	GitSubmodulesConfig Codebuild_ProjectSourceGitSubmodulesConfig `json:"gitSubmodulesConfig,omitempty" yaml:"gitSubmodulesConfig,omitempty"`

	// Ignore SSL warnings when connecting to source control.
	InsecureSsl bool `json:"insecureSsl,omitempty" yaml:"insecureSsl,omitempty"`

	// Location of the source code from git or s3.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Whether to report the status of a build's start and finish to your source provider. This option is valid only when your source provider is GitHub, GitHub Enterprise, GitLab, GitLab Self Managed, or Bitbucket.
	ReportBuildStatus bool `json:"reportBuildStatus,omitempty" yaml:"reportBuildStatus,omitempty"`

	// Type of repository that contains the source code to be built. Valid values: `BITBUCKET`, `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `GITLAB`, `GITLAB_SELF_MANAGED`, `NO_SOURCE`, `S3`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Configuration block that contains information that defines how the build project reports the build status to the source provider. This option is only used when the source provider is GitHub, GitHub Enterprise, GitLab, GitLab Self Managed, or Bitbucket. `build_status_config` blocks are documented below.
	BuildStatusConfig Codebuild_ProjectSourceBuildStatusConfig `json:"buildStatusConfig,omitempty" yaml:"buildStatusConfig,omitempty"`

	// Build specification to use for this build project's related builds. This must be set when `type` is `NO_SOURCE`. Also, if a non-default buildspec file name or file path aside from the root is used, it must be specified.
	Buildspec string `json:"buildspec,omitempty" yaml:"buildspec,omitempty"`

	// Truncate git history to this many commits. Use `0` for a `Full` checkout which you need to run commands like `git branch --show-current`. See [AWS CodePipeline User Guide: Tutorial: Use full clone with a GitHub pipeline source](https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-github-gitclone.html) for details.
	GitCloneDepth int `json:"gitCloneDepth,omitempty" yaml:"gitCloneDepth,omitempty"`
}
