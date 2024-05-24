package types

type Lb_ListenerDefaultActionAuthenticateOidc struct {
	// Authorization endpoint of the IdP.
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty" yaml:"authorizationEndpoint,omitempty"`

	// Behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" yaml:"onUnauthenticatedRequest,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`

	/*
	   User info endpoint of the IdP.

	   The following arguments are optional:
	*/
	UserInfoEndpoint string `json:"userInfoEndpoint,omitempty" yaml:"userInfoEndpoint,omitempty"`

	// Set of user claims to be requested from the IdP.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	SessionCookieName string `json:"sessionCookieName,omitempty" yaml:"sessionCookieName,omitempty"`

	// Token endpoint of the IdP.
	TokenEndpoint string `json:"tokenEndpoint,omitempty" yaml:"tokenEndpoint,omitempty"`

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	// OAuth 2.0 client identifier.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// OAuth 2.0 client secret.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// OIDC issuer identifier of the IdP.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
