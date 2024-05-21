package types



type Cloudbuildv2_ConnectionGithubEnterpriseConfig struct {
	// SecretManager resource containing the private key of the GitHub App, formatted as `projects/-/secrets/-/versions/-`.
	PrivateKeySecretVersion string `json:"privateKeySecretVersion,omitempty" yaml:"privateKeySecretVersion,omitempty"`

	/*
	   Configuration for using Service Directory to privately connect to a GitHub Enterprise server. This should only be set if the GitHub Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the GitHub Enterprise server will be made over the public internet.
	   Structure is documented below.
	*/
	ServiceDirectoryConfig Cloudbuildv2_ConnectionGithubEnterpriseConfigServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty" yaml:"serviceDirectoryConfig,omitempty"`

	// SSL certificate to use for requests to GitHub Enterprise.
	SslCa string `json:"sslCa,omitempty" yaml:"sslCa,omitempty"`

	// SecretManager resource containing the webhook secret of the GitHub App, formatted as `projects/-/secrets/-/versions/-`.
	WebhookSecretSecretVersion string `json:"webhookSecretSecretVersion,omitempty" yaml:"webhookSecretSecretVersion,omitempty"`

	// Id of the GitHub App created from the manifest.
	AppId int `json:"appId,omitempty" yaml:"appId,omitempty"`

	// ID of the installation of the GitHub App.
	AppInstallationId int `json:"appInstallationId,omitempty" yaml:"appInstallationId,omitempty"`

	// The URL-friendly name of the GitHub App.
	AppSlug string `json:"appSlug,omitempty" yaml:"appSlug,omitempty"`

	// Required. The URI of the GitHub Enterprise host this connection is for.
	HostUri string `json:"hostUri,omitempty" yaml:"hostUri,omitempty"`
}
