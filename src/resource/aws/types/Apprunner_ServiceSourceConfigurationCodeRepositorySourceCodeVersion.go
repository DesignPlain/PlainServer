package types

type Apprunner_ServiceSourceConfigurationCodeRepositorySourceCodeVersion struct {
	// Type of version identifier. For a git-based repository, branches represent versions. Valid values: `BRANCH`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
