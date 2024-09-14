package types

type Cloudbuildv2_ConnectionGitlabConfig struct {
	// Required. Immutable. SecretManager resource containing the webhook secret of a GitLab Enterprise project, formatted as `projects/-/secrets/-/versions/-`.
	WebhookSecretSecretVersion string `json:"webhookSecretSecretVersion,omitempty" yaml:"webhookSecretSecretVersion,omitempty"`

	/*
	   Required. A GitLab personal access token with the `api` scope access.
	   Structure is documented below.
	*/
	AuthorizerCredential Cloudbuildv2_ConnectionGitlabConfigAuthorizerCredential `json:"authorizerCredential,omitempty" yaml:"authorizerCredential,omitempty"`

	// The URI of the GitLab Enterprise host this connection is for. If not specified, the default value is https://gitlab.com.
	HostUri string `json:"hostUri,omitempty" yaml:"hostUri,omitempty"`

	/*
	   Required. A GitLab personal access token with the minimum `read_api` scope access.
	   Structure is documented below.
	*/
	ReadAuthorizerCredential Cloudbuildv2_ConnectionGitlabConfigReadAuthorizerCredential `json:"readAuthorizerCredential,omitempty" yaml:"readAuthorizerCredential,omitempty"`

	/*
	   (Output)
	   Output only. Version of the GitLab Enterprise server running on the `host_uri`.
	*/
	ServerVersion string `json:"serverVersion,omitempty" yaml:"serverVersion,omitempty"`

	/*
	   Configuration for using Service Directory to privately connect to a GitLab Enterprise server. This should only be set if the GitLab Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the GitLab Enterprise server will be made over the public internet.
	   Structure is documented below.
	*/
	ServiceDirectoryConfig Cloudbuildv2_ConnectionGitlabConfigServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty" yaml:"serviceDirectoryConfig,omitempty"`

	// SSL certificate to use for requests to GitLab Enterprise.
	SslCa string `json:"sslCa,omitempty" yaml:"sslCa,omitempty"`
}
