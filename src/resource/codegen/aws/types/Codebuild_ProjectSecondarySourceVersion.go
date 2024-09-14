package types

type Codebuild_ProjectSecondarySourceVersion struct {
	// The source version for the corresponding source identifier. See [AWS docs](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_ProjectSourceVersion.html#CodeBuild-Type-ProjectSourceVersion-sourceVersion) for more details.
	SourceVersion string `json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`

	// An identifier for a source in the build project.
	SourceIdentifier string `json:"sourceIdentifier,omitempty" yaml:"sourceIdentifier,omitempty"`
}
