package types

type Verifiedaccess_TrustProviderOidcOptions struct {
	//
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty" yaml:"authorizationEndpoint,omitempty"`

	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	//
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	//
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	//
	TokenEndpoint string `json:"tokenEndpoint,omitempty" yaml:"tokenEndpoint,omitempty"`

	//
	UserInfoEndpoint string `json:"userInfoEndpoint,omitempty" yaml:"userInfoEndpoint,omitempty"`
}
