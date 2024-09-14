package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigNpmRepository struct {
	// Address of the remote repository. Default value: "NPMJS" Possible values: ["NPMJS"]
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
