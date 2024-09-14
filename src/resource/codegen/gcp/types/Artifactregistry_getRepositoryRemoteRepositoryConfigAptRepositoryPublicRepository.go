package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository struct {
	// Specific repository from the base.
	RepositoryPath string `json:"repositoryPath,omitempty" yaml:"repositoryPath,omitempty"`

	// A common public repository base for Apt, e.g. '"debian/dists/buster"' Possible values: ["DEBIAN", "UBUNTU"]
	RepositoryBase string `json:"repositoryBase,omitempty" yaml:"repositoryBase,omitempty"`
}
