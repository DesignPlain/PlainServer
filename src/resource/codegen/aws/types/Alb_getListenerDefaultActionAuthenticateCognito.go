package types

type Alb_getListenerDefaultActionAuthenticateCognito struct {
	//
	UserPoolArn string `json:"userPoolArn,omitempty" yaml:"userPoolArn,omitempty"`

	//
	UserPoolClientId string `json:"userPoolClientId,omitempty" yaml:"userPoolClientId,omitempty"`

	//
	UserPoolDomain string `json:"userPoolDomain,omitempty" yaml:"userPoolDomain,omitempty"`

	//
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	//
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" yaml:"onUnauthenticatedRequest,omitempty"`

	//
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	//
	SessionCookieName string `json:"sessionCookieName,omitempty" yaml:"sessionCookieName,omitempty"`

	//
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`
}
