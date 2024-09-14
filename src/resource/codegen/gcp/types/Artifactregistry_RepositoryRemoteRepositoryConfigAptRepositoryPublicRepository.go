package types

type Artifactregistry_RepositoryRemoteRepositoryConfigAptRepositoryPublicRepository struct {
	// Specific repository from the base, e.g. `"centos/8-stream/BaseOS/x86_64/os"`
	RepositoryPath string `json:"repositoryPath,omitempty" yaml:"repositoryPath,omitempty"`

	/*
	   A common public repository base for Yum.
	   Possible values are: `CENTOS`, `CENTOS_DEBUG`, `CENTOS_VAULT`, `CENTOS_STREAM`, `ROCKY`, `EPEL`.
	*/
	RepositoryBase string `json:"repositoryBase,omitempty" yaml:"repositoryBase,omitempty"`
}
