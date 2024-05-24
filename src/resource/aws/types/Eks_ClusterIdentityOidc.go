package types

type Eks_ClusterIdentityOidc struct {
	// Issuer URL for the OpenID Connect identity provider.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
