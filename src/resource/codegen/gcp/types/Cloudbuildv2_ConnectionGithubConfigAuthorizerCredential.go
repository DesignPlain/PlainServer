package types

type Cloudbuildv2_ConnectionGithubConfigAuthorizerCredential struct {
	// A SecretManager resource containing the OAuth token that authorizes the Cloud Build connection. Format: `projects/-/secrets/-/versions/-`.
	OauthTokenSecretVersion string `json:"oauthTokenSecretVersion,omitempty" yaml:"oauthTokenSecretVersion,omitempty"`

	/*
	   (Output)
	   Output only. The username associated to this token.
	*/
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
