package identityplatform

type TenantDefaultSupportedIdpConfig struct {
	// OAuth client ID
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	/*
	   OAuth client secret


	   - - -
	*/
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// If this IDP allows the user to sign in
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   ID of the IDP. Possible values include:
	   - `apple.com`
	   - `facebook.com`
	   - `gc.apple.com`
	   - `github.com`
	   - `google.com`
	   - `linkedin.com`
	   - `microsoft.com`
	   - `playgames.google.com`
	   - `twitter.com`
	   - `yahoo.com`
	*/
	IdpId string `json:"idpId,omitempty" yaml:"idpId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the tenant where this DefaultSupportedIdpConfig resource exists
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`
}
