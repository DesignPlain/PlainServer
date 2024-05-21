package types

type Artifactregistry_RepositoryRemoteRepositoryConfigAptRepository struct {
	/*
	   One of the publicly available Apt repositories supported by Artifact Registry.
	   Structure is documented below.
	*/
	PublicRepository Artifactregistry_RepositoryRemoteRepositoryConfigAptRepositoryPublicRepository `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
