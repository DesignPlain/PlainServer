package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigYumRepository struct {
	// One of the publicly available Yum repositories supported by Artifact Registry.
	PublicRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository `json:"publicRepositories,omitempty" yaml:"publicRepositories,omitempty"`
}
