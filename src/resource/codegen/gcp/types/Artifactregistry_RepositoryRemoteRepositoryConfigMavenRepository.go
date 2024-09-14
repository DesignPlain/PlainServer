package types

type Artifactregistry_RepositoryRemoteRepositoryConfigMavenRepository struct {
	/*
	   Address of the remote repository.
	   Default value is `MAVEN_CENTRAL`.
	   Possible values are: `MAVEN_CENTRAL`.
	*/
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
