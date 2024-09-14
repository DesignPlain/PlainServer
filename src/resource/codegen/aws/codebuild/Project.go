package codebuild

import types "libds/aws/types"

type Project struct {
	// Configuration block. Detailed below.
	SecondaryArtifacts []types.Codebuild_ProjectSecondaryArtifact `json:"secondaryArtifacts,omitempty" yaml:"secondaryArtifacts,omitempty"`

	// Version of the build input to be built for this project. If not specified, the latest version is used.
	SourceVersion string `json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`

	// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
	ConcurrentBuildLimit int `json:"concurrentBuildLimit,omitempty" yaml:"concurrentBuildLimit,omitempty"`

	// A set of file system locations to mount inside the build. File system locations are documented below.
	FileSystemLocations []types.Codebuild_ProjectFileSystemLocation `json:"fileSystemLocations,omitempty" yaml:"fileSystemLocations,omitempty"`

	// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours. The `queued_timeout` property is not available on the `Lambda` compute type.
	QueuedTimeout int `json:"queuedTimeout,omitempty" yaml:"queuedTimeout,omitempty"`

	// Configuration block. Detailed below.
	LogsConfig types.Codebuild_ProjectLogsConfig `json:"logsConfig,omitempty" yaml:"logsConfig,omitempty"`

	// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
	ProjectVisibility string `json:"projectVisibility,omitempty" yaml:"projectVisibility,omitempty"`

	// Configuration block. Detailed below.
	SecondarySourceVersions []types.Codebuild_ProjectSecondarySourceVersion `json:"secondarySourceVersions,omitempty" yaml:"secondarySourceVersions,omitempty"`

	/*
	   Configuration block. Detailed below.

	   The following arguments are optional:
	*/
	Source types.Codebuild_ProjectSource `json:"source,omitempty" yaml:"source,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Generates a publicly-accessible URL for the projects build badge. Available as `badge_url` attribute when enabled.
	BadgeEnabled bool `json:"badgeEnabled,omitempty" yaml:"badgeEnabled,omitempty"`

	// Defines the batch build options for the project.
	BuildBatchConfig types.Codebuild_ProjectBuildBatchConfig `json:"buildBatchConfig,omitempty" yaml:"buildBatchConfig,omitempty"`

	// Configuration block. Detailed below.
	Cache types.Codebuild_ProjectCache `json:"cache,omitempty" yaml:"cache,omitempty"`

	// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`

	// Number of minutes, from 5 to 2160 (36 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes. The `build_timeout` property is not available on the `Lambda` compute type.
	BuildTimeout int `json:"buildTimeout,omitempty" yaml:"buildTimeout,omitempty"`

	// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
	EncryptionKey string `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`

	// Configuration block. Detailed below.
	Environment types.Codebuild_ProjectEnvironment `json:"environment,omitempty" yaml:"environment,omitempty"`

	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds in order to display them publicly. Only applicable if `project_visibility` is `PUBLIC_READ`.
	ResourceAccessRole string `json:"resourceAccessRole,omitempty" yaml:"resourceAccessRole,omitempty"`

	// Configuration block. Detailed below.
	SecondarySources []types.Codebuild_ProjectSecondarySource `json:"secondarySources,omitempty" yaml:"secondarySources,omitempty"`

	// Configuration block. Detailed below.
	VpcConfig types.Codebuild_ProjectVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Configuration block. Detailed below.
	Artifacts types.Codebuild_ProjectArtifacts `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`

	// Short description of the project.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Project's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
