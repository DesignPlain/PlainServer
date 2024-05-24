package types

type Lb_ListenerDefaultActionAuthenticateCognito struct {
	// Name of the cookie used to maintain session information.
	SessionCookieName string `json:"sessionCookieName,omitempty" yaml:"sessionCookieName,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`

	// ARN of the Cognito user pool.
	UserPoolArn string `json:"userPoolArn,omitempty" yaml:"userPoolArn,omitempty"`

	// ID of the Cognito user pool client.
	UserPoolClientId string `json:"userPoolClientId,omitempty" yaml:"userPoolClientId,omitempty"`

	/*
	   Domain prefix or fully-qualified domain name of the Cognito user pool.

	   The following arguments are optional:
	*/
	UserPoolDomain string `json:"userPoolDomain,omitempty" yaml:"userPoolDomain,omitempty"`

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10. Detailed below.
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" yaml:"authenticationRequestExtraParams,omitempty"`

	// Behavior if the user is not authenticated. Valid values are `deny`, `allow` and `authenticate`.
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" yaml:"onUnauthenticatedRequest,omitempty"`

	// Set of user claims to be requested from the IdP.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
