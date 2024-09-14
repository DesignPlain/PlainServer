package types

type Codegurureviewer_RepositoryAssociationRepositoryGithubEnterpriseServer struct {
	// The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
	ConnectionArn string `json:"connectionArn,omitempty" yaml:"connectionArn,omitempty"`

	// The name of the third party source repository.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The username for the account that owns the repository.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`
}
