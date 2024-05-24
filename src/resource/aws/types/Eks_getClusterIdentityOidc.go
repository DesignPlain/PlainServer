package types

type Eks_getClusterIdentityOidc struct {
	// Issuer URL for the OpenID Connect identity provider.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
