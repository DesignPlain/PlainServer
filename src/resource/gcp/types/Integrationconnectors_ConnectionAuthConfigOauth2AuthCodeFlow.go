package types

type Integrationconnectors_ConnectionAuthConfigOauth2AuthCodeFlow struct {
	// Whether to enable PKCE when the user performs the auth code flow.
	EnablePkce bool `json:"enablePkce,omitempty" yaml:"enablePkce,omitempty"`

	// Scopes the connection will request when the user performs the auth code flow.
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`

	// Auth URL for Authorization Code Flow.
	AuthUri string `json:"authUri,omitempty" yaml:"authUri,omitempty"`

	// Secret version of Password for Authentication.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	/*
	   Secret version reference containing the client secret.
	   Structure is documented below.
	*/
	ClientSecret Integrationconnectors_ConnectionAuthConfigOauth2AuthCodeFlowClientSecret `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
}
