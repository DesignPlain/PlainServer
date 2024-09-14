package cognito

import types "libds/aws/types"

type ManagedUserPoolClient struct {
	// Regular expression that matches the name of the desired User Pool Client. It must only match one User Pool Client.
	NamePattern string `json:"namePattern,omitempty" yaml:"namePattern,omitempty"`

	// Time limit, between 60 minutes and 10 years, after which the refresh token is no longer valid and cannot be used. By default, the unit is days. The unit can be overridden by a value in `token_validity_units.refresh_token`.
	RefreshTokenValidity int `json:"refreshTokenValidity,omitempty" yaml:"refreshTokenValidity,omitempty"`

	// Configuration block for representing the validity times in units. See details below. Detailed below.
	TokenValidityUnits types.Cognito_ManagedUserPoolClientTokenValidityUnits `json:"tokenValidityUnits,omitempty" yaml:"tokenValidityUnits,omitempty"`

	// Whether the client is allowed to use OAuth 2.0 features. `allowed_oauth_flows_user_pool_client` must be set to `true` before you can configure the following arguments: `callback_urls`, `logout_urls`, `allowed_oauth_scopes` and `allowed_oauth_flows`.
	AllowedOauthFlowsUserPoolClient bool `json:"allowedOauthFlowsUserPoolClient,omitempty" yaml:"allowedOauthFlowsUserPoolClient,omitempty"`

	// List of allowed OAuth flows, including `code`, `implicit`, and `client_credentials`. `allowed_oauth_flows_user_pool_client` must be set to `true` before you can configure this option.
	AllowedOauthFlows []string `json:"allowedOauthFlows,omitempty" yaml:"allowedOauthFlows,omitempty"`

	// List of allowed OAuth scopes, including `phone`, `email`, `openid`, `profile`, and `aws.cognito.signin.user.admin`. `allowed_oauth_flows_user_pool_client` must be set to `true` before you can configure this option.
	AllowedOauthScopes []string `json:"allowedOauthScopes,omitempty" yaml:"allowedOauthScopes,omitempty"`

	// Configuration block for Amazon Pinpoint analytics that collects metrics for this user pool. See details below.
	AnalyticsConfiguration types.Cognito_ManagedUserPoolClientAnalyticsConfiguration `json:"analyticsConfiguration,omitempty" yaml:"analyticsConfiguration,omitempty"`

	// Duration, in minutes, of the session token created by Amazon Cognito for each API request in an authentication flow. The session token must be responded to by the native user of the user pool before it expires. Valid values for `auth_session_validity` are between `3` and `15`, with a default value of `3`.
	AuthSessionValidity int `json:"authSessionValidity,omitempty" yaml:"authSessionValidity,omitempty"`

	// List of allowed callback URLs for the identity providers. `allowed_oauth_flows_user_pool_client` must be set to `true` before you can configure this option.
	CallbackUrls []string `json:"callbackUrls,omitempty" yaml:"callbackUrls,omitempty"`

	// Enables the propagation of additional user context data.
	EnablePropagateAdditionalUserContextData bool `json:"enablePropagateAdditionalUserContextData,omitempty" yaml:"enablePropagateAdditionalUserContextData,omitempty"`

	// Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used. By default, the unit is hours. The unit can be overridden by a value in `token_validity_units.id_token`.
	IdTokenValidity int `json:"idTokenValidity,omitempty" yaml:"idTokenValidity,omitempty"`

	// Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used. By default, the unit is hours. The unit can be overridden by a value in `token_validity_units.access_token`.
	AccessTokenValidity int `json:"accessTokenValidity,omitempty" yaml:"accessTokenValidity,omitempty"`

	// List of user pool attributes that the application client can read from.
	ReadAttributes []string `json:"readAttributes,omitempty" yaml:"readAttributes,omitempty"`

	// Setting determines the errors and responses returned by Cognito APIs when a user does not exist in the user pool during authentication, account confirmation, and password recovery.
	PreventUserExistenceErrors string `json:"preventUserExistenceErrors,omitempty" yaml:"preventUserExistenceErrors,omitempty"`

	// List of user pool attributes that the application client can write to.
	WriteAttributes []string `json:"writeAttributes,omitempty" yaml:"writeAttributes,omitempty"`

	// List of authentication flows. The available options include ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, and ALLOW_REFRESH_TOKEN_AUTH.
	ExplicitAuthFlows []string `json:"explicitAuthFlows,omitempty" yaml:"explicitAuthFlows,omitempty"`

	// Enables or disables token revocation.
	EnableTokenRevocation bool `json:"enableTokenRevocation,omitempty" yaml:"enableTokenRevocation,omitempty"`

	// List of allowed logout URLs for the identity providers. `allowed_oauth_flows_user_pool_client` must be set to `true` before you can configure this option.
	LogoutUrls []string `json:"logoutUrls,omitempty" yaml:"logoutUrls,omitempty"`

	/*
	   String that matches the beginning of the name of the desired User Pool Client. It must match only one User Pool Client.

	   The following arguments are optional:
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// List of provider names for the identity providers that are supported on this client. It uses the `provider_name` attribute of the `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
	SupportedIdentityProviders []string `json:"supportedIdentityProviders,omitempty" yaml:"supportedIdentityProviders,omitempty"`

	// User pool that the client belongs to.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// Default redirect URI and must be included in the list of callback URLs.
	DefaultRedirectUri string `json:"defaultRedirectUri,omitempty" yaml:"defaultRedirectUri,omitempty"`
}
