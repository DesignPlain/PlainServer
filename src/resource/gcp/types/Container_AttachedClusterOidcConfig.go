package types

type Container_AttachedClusterOidcConfig struct {
	// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://`
	IssuerUrl string `json:"issuerUrl,omitempty" yaml:"issuerUrl,omitempty"`

	// OIDC verification keys in JWKS format (RFC 7517).
	Jwks string `json:"jwks,omitempty" yaml:"jwks,omitempty"`
}
