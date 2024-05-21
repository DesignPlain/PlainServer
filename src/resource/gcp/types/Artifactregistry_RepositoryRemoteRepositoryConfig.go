package types

type Artifactregistry_RepositoryRemoteRepositoryConfig struct {
	/*
	   Specific settings for an Apt remote repository.
	   Structure is documented below.
	*/
	AptRepository Artifactregistry_RepositoryRemoteRepositoryConfigAptRepository `json:"aptRepository,omitempty" yaml:"aptRepository,omitempty"`

	// The description of the remote source.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specific settings for a Docker remote repository.
	   Structure is documented below.
	*/
	DockerRepository Artifactregistry_RepositoryRemoteRepositoryConfigDockerRepository `json:"dockerRepository,omitempty" yaml:"dockerRepository,omitempty"`

	/*
	   Specific settings for a Maven remote repository.
	   Structure is documented below.
	*/
	MavenRepository Artifactregistry_RepositoryRemoteRepositoryConfigMavenRepository `json:"mavenRepository,omitempty" yaml:"mavenRepository,omitempty"`

	/*
	   Specific settings for an Npm remote repository.
	   Structure is documented below.
	*/
	NpmRepository Artifactregistry_RepositoryRemoteRepositoryConfigNpmRepository `json:"npmRepository,omitempty" yaml:"npmRepository,omitempty"`

	/*
	   Specific settings for a Python remote repository.
	   Structure is documented below.
	*/
	PythonRepository Artifactregistry_RepositoryRemoteRepositoryConfigPythonRepository `json:"pythonRepository,omitempty" yaml:"pythonRepository,omitempty"`

	/*
	   The credentials used to access the remote repository.
	   Structure is documented below.
	*/
	UpstreamCredentials Artifactregistry_RepositoryRemoteRepositoryConfigUpstreamCredentials `json:"upstreamCredentials,omitempty" yaml:"upstreamCredentials,omitempty"`

	/*
	   Specific settings for an Yum remote repository.
	   Structure is documented below.
	*/
	YumRepository Artifactregistry_RepositoryRemoteRepositoryConfigYumRepository `json:"yumRepository,omitempty" yaml:"yumRepository,omitempty"`
}
