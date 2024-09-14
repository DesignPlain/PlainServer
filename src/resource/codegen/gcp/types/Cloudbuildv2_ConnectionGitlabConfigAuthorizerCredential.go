package types

type Cloudbuildv2_ConnectionGitlabConfigAuthorizerCredential struct {
	// Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: `projects/-/secrets/-/versions/-`.
	UserTokenSecretVersion string `json:"userTokenSecretVersion,omitempty" yaml:"userTokenSecretVersion,omitempty"`

	/*
	   (Output)
	   Output only. The username associated to this token.
	*/
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
