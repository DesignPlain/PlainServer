package types

type Artifactregistry_getRepositoryRemoteRepositoryConfigPythonRepository struct {
	// Address of the remote repository. Default value: "PYPI" Possible values: ["PYPI"]
	PublicRepository string `json:"publicRepository,omitempty" yaml:"publicRepository,omitempty"`
}
