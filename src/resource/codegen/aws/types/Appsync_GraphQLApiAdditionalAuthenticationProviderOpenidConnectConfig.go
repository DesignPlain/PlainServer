package types

type Appsync_GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig struct {
	// Number of milliseconds a token is valid after being authenticated.
	AuthTtl int `json:"authTtl,omitempty" yaml:"authTtl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	IatTtl int `json:"iatTtl,omitempty" yaml:"iatTtl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
