package types

type Appfabric_AppAuthorizationConnectionAuthRequest struct {
	// The authorization code returned by the application after permission is granted in the application OAuth page (after clicking on the AuthURL)..
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// The redirect URL that is specified in the AuthURL and the application client.
	RedirectUri string `json:"redirectUri,omitempty" yaml:"redirectUri,omitempty"`
}
