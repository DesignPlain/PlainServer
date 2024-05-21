package types

type Integrationconnectors_ConnectionAuthConfigOauth2ClientCredentials struct {
	// Secret version of Password for Authentication.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	/*
	   Secret version reference containing the client secret.
	   Structure is documented below.
	*/
	ClientSecret Integrationconnectors_ConnectionAuthConfigOauth2ClientCredentialsClientSecret `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
}
