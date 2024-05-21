package types

type Artifactregistry_RepositoryRemoteRepositoryConfigDockerRepository struct {
	/*
	   Address of the remote repository.
	   Default value is `DOCKER_HUB`.
	   Possible values are: `DOCKER_HUB`.
	*/
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
