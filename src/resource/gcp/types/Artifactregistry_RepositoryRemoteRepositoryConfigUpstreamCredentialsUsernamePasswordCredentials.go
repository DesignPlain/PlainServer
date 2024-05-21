package types

type Artifactregistry_RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials struct {
	/*
	   The Secret Manager key version that holds the password to access the
	   remote repository. Must be in the format of
	   `projects/{project}/secrets/{secret}/versions/{version}`.
	*/
	PasswordSecretVersion string `json:"passwordSecretVersion,omitempty" yaml:"passwordSecretVersion,omitempty"`

	// The username to access the remote repository.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
