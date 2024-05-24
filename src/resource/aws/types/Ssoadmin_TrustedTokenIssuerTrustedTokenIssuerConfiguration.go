package types

type Ssoadmin_TrustedTokenIssuerTrustedTokenIssuerConfiguration struct {
	// A block that describes the settings for a trusted token issuer that works with OpenID Connect (OIDC) by using JSON Web Tokens (JWT). See Documented below below.
	OidcJwtConfiguration Ssoadmin_TrustedTokenIssuerTrustedTokenIssuerConfigurationOidcJwtConfiguration `json:"oidcJwtConfiguration,omitempty" yaml:"oidcJwtConfiguration,omitempty"`
}
