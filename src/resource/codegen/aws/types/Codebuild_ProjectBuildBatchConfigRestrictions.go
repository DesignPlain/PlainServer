package types

type Codebuild_ProjectBuildBatchConfigRestrictions struct {
	// Specifies the maximum number of builds allowed.
	MaximumBuildsAllowed int `json:"maximumBuildsAllowed,omitempty" yaml:"maximumBuildsAllowed,omitempty"`

	// An array of strings that specify the compute types that are allowed for the batch build. See [Build environment compute types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the AWS CodeBuild User Guide for these values.
	ComputeTypesAlloweds []string `json:"computeTypesAlloweds,omitempty" yaml:"computeTypesAlloweds,omitempty"`
}
