package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigMavenRepository struct {
	// Address of the remote repository. Default value: "MAVEN_CENTRAL" Possible values: ["MAVEN_CENTRAL"]
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
