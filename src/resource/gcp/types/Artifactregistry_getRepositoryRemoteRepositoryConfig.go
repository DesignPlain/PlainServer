package types

type Artifactregistry_getRepositoryRemoteRepositoryConfig struct {
	// Specific settings for an Npm remote repository.
	NpmRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigNpmRepository `json:"npmRepositories,omitempty" yaml:"npmRepositories,omitempty"`

	// Specific settings for a Python remote repository.
	PythonRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigPythonRepository `json:"pythonRepositories,omitempty" yaml:"pythonRepositories,omitempty"`

	// The credentials used to access the remote repository.
	UpstreamCredentials []Artifactregistry_getRepositoryRemoteRepositoryConfigUpstreamCredential `json:"upstreamCredentials,omitempty" yaml:"upstreamCredentials,omitempty"`

	// Specific settings for an Yum remote repository.
	YumRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigYumRepository `json:"yumRepositories,omitempty" yaml:"yumRepositories,omitempty"`

	// Specific settings for an Apt remote repository.
	AptRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigAptRepository `json:"aptRepositories,omitempty" yaml:"aptRepositories,omitempty"`

	// The description of the remote source.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specific settings for a Docker remote repository.
	DockerRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigDockerRepository `json:"dockerRepositories,omitempty" yaml:"dockerRepositories,omitempty"`

	// Specific settings for a Maven remote repository.
	MavenRepositories []Artifactregistry_getRepositoryRemoteRepositoryConfigMavenRepository `json:"mavenRepositories,omitempty" yaml:"mavenRepositories,omitempty"`
}
