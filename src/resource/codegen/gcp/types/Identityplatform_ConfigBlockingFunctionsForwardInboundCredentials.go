package types

type Identityplatform_ConfigBlockingFunctionsForwardInboundCredentials struct {
	// Whether to pass the user's OAuth identity provider's access token.
	AccessToken bool `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// Whether to pass the user's OIDC identity provider's ID token.
	IdToken bool `json:"idToken,omitempty" yaml:"idToken,omitempty"`

	// Whether to pass the user's OAuth identity provider's refresh token.
	RefreshToken bool `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`
}
