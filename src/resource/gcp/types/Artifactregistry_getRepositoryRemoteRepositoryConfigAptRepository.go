package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigAptRepository struct {
	// One of the publicly available Apt repositories supported by Artifact Registry.
	PublicRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository `json:"publicRepositories,omitempty" yaml:"publicRepositories,omitempty"`
}
