package types

type Codecatalyst_DevEnvironmentRepository struct {
	/*
	   The name of the branch in a source repository.

	   persistent storage (` persistent_storage`) supports the following:
	*/
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// The name of the source repository.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
