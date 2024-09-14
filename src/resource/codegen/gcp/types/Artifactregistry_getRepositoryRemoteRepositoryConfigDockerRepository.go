package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigDockerRepository struct {
	// Address of the remote repository. Default value: "DOCKER_HUB" Possible values: ["DOCKER_HUB"]
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
