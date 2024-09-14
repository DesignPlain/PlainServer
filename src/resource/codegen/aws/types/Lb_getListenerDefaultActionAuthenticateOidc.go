package types

type Lb_getListenerDefaultActionAuthenticateOidc struct {
	//
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	//
	SessionCookieName string `json:"sessionCookieName,omitempty" yaml:"sessionCookieName,omitempty"`

	//
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`

	//
	TokenEndpoint string `json:"tokenEndpoint,omitempty" yaml:"tokenEndpoint,omitempty"`

	//
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty" yaml:"authorizationEndpoint,omitempty"`

	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" yaml:"onUnauthenticatedRequest,omitempty"`

	//
	UserInfoEndpoint string `json:"userInfoEndpoint,omitempty" yaml:"userInfoEndpoint,omitempty"`

	//
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	//
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	//
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
