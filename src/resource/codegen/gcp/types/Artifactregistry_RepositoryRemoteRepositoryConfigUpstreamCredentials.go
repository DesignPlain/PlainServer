package types

type Artifactregistry_RepositoryRemoteRepositoryConfigUpstreamCredentials struct {
	/*
	   Use username and password to access the remote repository.
	   Structure is documented below.
	*/
	UsernamePasswordCredentials Artifactregistry_RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials `json:"usernamePasswordCredentials,omitempty" yaml:"usernamePasswordCredentials,omitempty"`
}
