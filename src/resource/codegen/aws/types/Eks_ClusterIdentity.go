package types

type Eks_ClusterIdentity struct {
	// Nested block containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster. Detailed below.
	Oidcs []Eks_ClusterIdentityOidc `json:"oidcs,omitempty" yaml:"oidcs,omitempty"`
}
