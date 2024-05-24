package types

type Codebuild_ProjectEnvironmentEnvironmentVariable struct {
	// Project's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Build output artifact's type. Valid values: `CODEPIPELINE`, `NO_ARTIFACTS`, `S3`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Environment variable's value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
