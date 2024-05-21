package types

type Artifactregistry_RepositoryRemoteRepositoryConfigNpmRepository struct {
	/*
	   Address of the remote repository.
	   Default value is `NPMJS`.
	   Possible values are: `NPMJS`.
	*/
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
