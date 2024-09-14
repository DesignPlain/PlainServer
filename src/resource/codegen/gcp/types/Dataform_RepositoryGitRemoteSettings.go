package types

type Dataform_RepositoryGitRemoteSettings struct {
	// The name of the Secret Manager secret version to use as an authentication token for Git operations. This secret is for assigning with HTTPS only(for SSH use `ssh_authentication_config`). Must be in the format projects/-/secrets/-/versions/-.
	AuthenticationTokenSecretVersion string `json:"authenticationTokenSecretVersion,omitempty" yaml:"authenticationTokenSecretVersion,omitempty"`

	// The Git remote's default branch name.
	DefaultBranch string `json:"defaultBranch,omitempty" yaml:"defaultBranch,omitempty"`

	/*
	   Authentication fields for remote uris using SSH protocol.
	   Structure is documented below.
	*/
	SshAuthenticationConfig Dataform_RepositoryGitRemoteSettingsSshAuthenticationConfig `json:"sshAuthenticationConfig,omitempty" yaml:"sshAuthenticationConfig,omitempty"`

	/*
	   (Output)
	   Indicates the status of the Git access token. https://cloud.google.com/dataform/reference/rest/v1beta1/projects.locations.repositories#TokenStatus
	*/
	TokenStatus string `json:"tokenStatus,omitempty" yaml:"tokenStatus,omitempty"`

	// The Git remote's URL.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
