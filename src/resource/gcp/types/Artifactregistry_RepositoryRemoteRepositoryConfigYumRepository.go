package types

type Artifactregistry_RepositoryRemoteRepositoryConfigYumRepository struct {
	/*
	   One of the publicly available Yum repositories supported by Artifact Registry.
	   Structure is documented below.
	*/
	PublicRepository Artifactregistry_RepositoryRemoteRepositoryConfigYumRepositoryPublicRepository `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
