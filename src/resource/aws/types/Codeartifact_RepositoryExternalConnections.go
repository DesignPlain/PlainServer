package types

type Codeartifact_RepositoryExternalConnections struct {
	// The name of the external connection associated with a repository.
	ExternalConnectionName string `json:"externalConnectionName,omitempty" yaml:"externalConnectionName,omitempty"`

	//
	PackageFormat string `json:"packageFormat,omitempty" yaml:"packageFormat,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
