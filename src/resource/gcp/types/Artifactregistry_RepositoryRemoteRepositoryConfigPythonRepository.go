package types

type Artifactregistry_RepositoryRemoteRepositoryConfigPythonRepository struct {
	/*
	   Address of the remote repository.
	   Default value is `PYPI`.
	   Possible values are: `PYPI`.
	*/
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
