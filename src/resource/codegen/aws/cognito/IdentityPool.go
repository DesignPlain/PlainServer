package cognito

import types "libds/aws/types"

type IdentityPool struct {
	/*
	   The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	   backend and the Cognito service to communicate about the developer provider.
	*/
	DeveloperProviderName string `json:"developerProviderName,omitempty" yaml:"developerProviderName,omitempty"`

	// A map of tags to assign to the Identity Pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Set of OpendID Connect provider ARNs.
	OpenidConnectProviderArns []string `json:"openidConnectProviderArns,omitempty" yaml:"openidConnectProviderArns,omitempty"`

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SamlProviderArns []string `json:"samlProviderArns,omitempty" yaml:"samlProviderArns,omitempty"`

	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string `json:"supportedLoginProviders,omitempty" yaml:"supportedLoginProviders,omitempty"`

	// Enables or disables the classic / basic authentication flow. Default is `false`.
	AllowClassicFlow bool `json:"allowClassicFlow,omitempty" yaml:"allowClassicFlow,omitempty"`

	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities bool `json:"allowUnauthenticatedIdentities,omitempty" yaml:"allowUnauthenticatedIdentities,omitempty"`

	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders []types.Cognito_IdentityPoolCognitoIdentityProvider `json:"cognitoIdentityProviders,omitempty" yaml:"cognitoIdentityProviders,omitempty"`

	// The Cognito Identity Pool name.
	IdentityPoolName string `json:"identityPoolName,omitempty" yaml:"identityPoolName,omitempty"`
}
