package identityplatform

type TenantOauthIdpConfig struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`

	/*
	   The client id of an OAuth client.


	   - - -
	*/
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// Human friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// For OIDC Idps, the issuer identifier.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with `oidc.`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
