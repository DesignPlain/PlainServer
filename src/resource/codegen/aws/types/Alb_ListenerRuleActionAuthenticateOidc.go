package types

type Alb_ListenerRuleActionAuthenticateOidc struct {
	// The set of user claims to be requested from the IdP.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// The name of the cookie used to maintain session information.
	SessionCookieName string `json:"sessionCookieName,omitempty" yaml:"sessionCookieName,omitempty"`

	// The OAuth 2.0 client identifier.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The OIDC issuer identifier of the IdP.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// The OAuth 2.0 client secret.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" yaml:"onUnauthenticatedRequest,omitempty"`

	// The maximum duration of the authentication session, in seconds.
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`

	// The token endpoint of the IdP.
	TokenEndpoint string `json:"tokenEndpoint,omitempty" yaml:"tokenEndpoint,omitempty"`

	// The user info endpoint of the IdP.
	UserInfoEndpoint string `json:"userInfoEndpoint,omitempty" yaml:"userInfoEndpoint,omitempty"`

	// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	// The authorization endpoint of the IdP.
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty" yaml:"authorizationEndpoint,omitempty"`
}
