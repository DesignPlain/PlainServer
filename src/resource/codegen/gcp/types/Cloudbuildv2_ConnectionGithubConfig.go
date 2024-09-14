package types

type Cloudbuildv2_ConnectionGithubConfig struct {
	// GitHub App installation id.
	AppInstallationId int `json:"appInstallationId,omitempty" yaml:"appInstallationId,omitempty"`

	/*
	   OAuth credential of the account that authorized the Cloud Build GitHub App. It is recommended to use a robot account instead of a human user account. The OAuth token must be tied to the Cloud Build GitHub App.
	   Structure is documented below.
	*/
	AuthorizerCredential Cloudbuildv2_ConnectionGithubConfigAuthorizerCredential `json:"authorizerCredential,omitempty" yaml:"authorizerCredential,omitempty"`
}
