package types

type Iam_WorkforcePoolProviderOidcWebSsoConfig struct {
	/*
	   Additional scopes to request for in the OIDC authentication request on top of scopes requested by default. By default, the `openid`, `profile` and `email` scopes that are supported by the identity provider are requested.
	   Each additional scope may be at most 256 characters. A maximum of 10 additional scopes may be configured.
	*/
	AdditionalScopes []string `json:"additionalScopes,omitempty" yaml:"additionalScopes,omitempty"`

	/*
	   The behavior for how OIDC Claims are included in the `assertion` object used for attribute mapping and attribute condition.
	   - MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS: Merge the UserInfo Endpoint Claims with ID Token Claims, preferring UserInfo Claim Values for the same Claim Name. This option is available only for the Authorization Code Flow.
	   - ONLY_ID_TOKEN_CLAIMS: Only include ID Token Claims.
	   Possible values are: `MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS`, `ONLY_ID_TOKEN_CLAIMS`.
	*/
	AssertionClaimsBehavior string `json:"assertionClaimsBehavior,omitempty" yaml:"assertionClaimsBehavior,omitempty"`

	/*
	   The Response Type to request for in the OIDC Authorization Request for web sign-in.
	   The `CODE` Response Type is recommended to avoid the Implicit Flow, for security reasons.
	   - CODE: The `response_type=code` selection uses the Authorization Code Flow for web sign-in. Requires a configured client secret.
	   - ID_TOKEN: The `response_type=id_token` selection uses the Implicit Flow for web sign-in.
	   Possible values are: `CODE`, `ID_TOKEN`.
	*/
	ResponseType string `json:"responseType,omitempty" yaml:"responseType,omitempty"`
}
