package types

type Sagemaker_WorkforceOidcConfig struct {
	// The OIDC IdP client secret used to configure your private workforce.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// The OIDC IdP logout endpoint used to configure your private workforce.
	LogoutEndpoint string `json:"logoutEndpoint,omitempty" yaml:"logoutEndpoint,omitempty"`

	// An array of string identifiers used to refer to the specific pieces of user data or claims that the client application wants to access.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// The OIDC IdP issuer used to configure your private workforce.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
	JwksUri string `json:"jwksUri,omitempty" yaml:"jwksUri,omitempty"`

	// The OIDC IdP token endpoint used to configure your private workforce.
	TokenEndpoint string `json:"tokenEndpoint,omitempty" yaml:"tokenEndpoint,omitempty"`

	// The OIDC IdP user information endpoint used to configure your private workforce.
	UserInfoEndpoint string `json:"userInfoEndpoint,omitempty" yaml:"userInfoEndpoint,omitempty"`

	// A string to string map of identifiers specific to the custom identity provider (IdP) being used.
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	// The OIDC IdP authorization endpoint used to configure your private workforce.
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty" yaml:"authorizationEndpoint,omitempty"`

	// The OIDC IdP client ID used to configure your private workforce.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
}
