package types

type Codebuild_ProjectBuildBatchConfig struct {
	// Specifies if the build artifacts for the batch build should be combined into a single artifact location.
	CombineArtifacts bool `json:"combineArtifacts,omitempty" yaml:"combineArtifacts,omitempty"`

	// Configuration block specifying the restrictions for the batch build. Detailed below.
	Restrictions Codebuild_ProjectBuildBatchConfigRestrictions `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`

	// Specifies the service role ARN for the batch build project.
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`

	// Specifies the maximum amount of time, in minutes, that the batch build must be completed in.
	TimeoutInMins int `json:"timeoutInMins,omitempty" yaml:"timeoutInMins,omitempty"`
}
