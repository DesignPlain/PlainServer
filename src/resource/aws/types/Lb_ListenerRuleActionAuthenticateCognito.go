package types

type Lb_ListenerRuleActionAuthenticateCognito struct {
	// The name of the cookie used to maintain session information.
	SessionCookieName string `json:"sessionCookieName,omitempty" yaml:"sessionCookieName,omitempty"`

	// The maximum duration of the authentication session, in seconds.
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`

	// The ARN of the Cognito user pool.
	UserPoolArn string `json:"userPoolArn,omitempty" yaml:"userPoolArn,omitempty"`

	// The ID of the Cognito user pool client.
	UserPoolClientId string `json:"userPoolClientId,omitempty" yaml:"userPoolClientId,omitempty"`

	// The domain prefix or fully-qualified domain name of the Cognito user pool.
	UserPoolDomain string `json:"userPoolDomain,omitempty" yaml:"userPoolDomain,omitempty"`

	// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" yaml:"onUnauthenticatedRequest,omitempty"`

	// The set of user claims to be requested from the IdP.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
