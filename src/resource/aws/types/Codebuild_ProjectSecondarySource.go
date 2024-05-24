package types

type Codebuild_ProjectSecondarySource struct {
	// An identifier for this project source. The identifier can only contain alphanumeric characters and underscores, and must be less than 128 characters in length.
	SourceIdentifier string `json:"sourceIdentifier,omitempty" yaml:"sourceIdentifier,omitempty"`

	// Configuration block that contains information that defines how the build project reports the build status to the source provider. This option is only used when the source provider is `GITHUB`, `GITHUB_ENTERPRISE`, or `BITBUCKET`. `build_status_config` blocks are documented below.
	BuildStatusConfig Codebuild_ProjectSecondarySourceBuildStatusConfig `json:"buildStatusConfig,omitempty" yaml:"buildStatusConfig,omitempty"`

	// The build spec declaration to use for this build project's related builds. This must be set when `type` is `NO_SOURCE`. It can either be a path to a file residing in the repository to be built or a local file path leveraging the `file()` built-in.
	Buildspec string `json:"buildspec,omitempty" yaml:"buildspec,omitempty"`

	// Truncate git history to this many commits. Use `0` for a `Full` checkout which you need to run commands like `git branch --show-current`. See [AWS CodePipeline User Guide: Tutorial: Use full clone with a GitHub pipeline source](https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-github-gitclone.html) for details.
	GitCloneDepth int `json:"gitCloneDepth,omitempty" yaml:"gitCloneDepth,omitempty"`

	// Ignore SSL warnings when connecting to source control.
	InsecureSsl bool `json:"insecureSsl,omitempty" yaml:"insecureSsl,omitempty"`

	// Location of the source code from git or s3.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Configuration block. Detailed below.
	GitSubmodulesConfig Codebuild_ProjectSecondarySourceGitSubmodulesConfig `json:"gitSubmodulesConfig,omitempty" yaml:"gitSubmodulesConfig,omitempty"`

	// Whether to report the status of a build's start and finish to your source provider. This option is only valid when your source provider is `GITHUB`, `BITBUCKET`, or `GITHUB_ENTERPRISE`.
	ReportBuildStatus bool `json:"reportBuildStatus,omitempty" yaml:"reportBuildStatus,omitempty"`

	// Type of repository that contains the source code to be built. Valid values: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
