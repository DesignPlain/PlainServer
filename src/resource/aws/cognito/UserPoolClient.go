package cognito

import types "DesignSphere_Server/src/resource/aws/types"

type UserPoolClient struct {
	// Name of the application client.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Time limit, between 60 minutes and 10 years, after which the refresh token is no longer valid and cannot be used.
	   By default, the unit is days.
	   The unit can be overridden by a value in `token_validity_units.refresh_token`.
	*/
	RefreshTokenValidity int `json:"refreshTokenValidity,omitempty" yaml:"refreshTokenValidity,omitempty"`

	/*
	   Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used.
	   By default, the unit is hours.
	   The unit can be overridden by a value in `token_validity_units.access_token`.
	*/
	AccessTokenValidity int `json:"accessTokenValidity,omitempty" yaml:"accessTokenValidity,omitempty"`

	// Enables or disables token revocation.
	EnableTokenRevocation bool `json:"enableTokenRevocation,omitempty" yaml:"enableTokenRevocation,omitempty"`

	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	ExplicitAuthFlows []string `json:"explicitAuthFlows,omitempty" yaml:"explicitAuthFlows,omitempty"`

	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors string `json:"preventUserExistenceErrors,omitempty" yaml:"preventUserExistenceErrors,omitempty"`

	// List of user pool attributes the application client can read from.
	ReadAttributes []string `json:"readAttributes,omitempty" yaml:"readAttributes,omitempty"`

	// List of user pool attributes the application client can write to.
	WriteAttributes []string `json:"writeAttributes,omitempty" yaml:"writeAttributes,omitempty"`

	/*
	   Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used.
	   By default, the unit is hours.
	   The unit can be overridden by a value in `token_validity_units.id_token`.
	*/
	IdTokenValidity int `json:"idTokenValidity,omitempty" yaml:"idTokenValidity,omitempty"`

	// List of allowed OAuth flows (code, implicit, client_credentials).
	AllowedOauthFlows []string `json:"allowedOauthFlows,omitempty" yaml:"allowedOauthFlows,omitempty"`

	// Configuration block for Amazon Pinpoint analytics for collecting metrics for this user pool. Detailed below.
	AnalyticsConfiguration types.Cognito_UserPoolClientAnalyticsConfiguration `json:"analyticsConfiguration,omitempty" yaml:"analyticsConfiguration,omitempty"`

	// Amazon Cognito creates a session token for each API request in an authentication flow. AuthSessionValidity is the duration, in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires. Valid values between `3` and `15`. Default value is `3`.
	AuthSessionValidity int `json:"authSessionValidity,omitempty" yaml:"authSessionValidity,omitempty"`

	// List of allowed callback URLs for the identity providers.
	CallbackUrls []string `json:"callbackUrls,omitempty" yaml:"callbackUrls,omitempty"`

	// Default redirect URI. Must be in the list of callback URLs.
	DefaultRedirectUri string `json:"defaultRedirectUri,omitempty" yaml:"defaultRedirectUri,omitempty"`

	// Activates the propagation of additional user context data.
	EnablePropagateAdditionalUserContextData bool `json:"enablePropagateAdditionalUserContextData,omitempty" yaml:"enablePropagateAdditionalUserContextData,omitempty"`

	// Should an application secret be generated.
	GenerateSecret bool `json:"generateSecret,omitempty" yaml:"generateSecret,omitempty"`

	// Configuration block for units in which the validity times are represented in. Detailed below.
	TokenValidityUnits types.Cognito_UserPoolClientTokenValidityUnits `json:"tokenValidityUnits,omitempty" yaml:"tokenValidityUnits,omitempty"`

	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	AllowedOauthFlowsUserPoolClient bool `json:"allowedOauthFlowsUserPoolClient,omitempty" yaml:"allowedOauthFlowsUserPoolClient,omitempty"`

	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	AllowedOauthScopes []string `json:"allowedOauthScopes,omitempty" yaml:"allowedOauthScopes,omitempty"`

	// List of allowed logout URLs for the identity providers.
	LogoutUrls []string `json:"logoutUrls,omitempty" yaml:"logoutUrls,omitempty"`

	// List of provider names for the identity providers that are supported on this client. Uses the `provider_name` attribute of `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
	SupportedIdentityProviders []string `json:"supportedIdentityProviders,omitempty" yaml:"supportedIdentityProviders,omitempty"`

	/*
	   User pool the client belongs to.

	   The following arguments are optional:
	*/
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
