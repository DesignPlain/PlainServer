package types

type Eks_getClusterIdentity struct {
	// Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster.
	Oidcs []Eks_getClusterIdentityOidc `json:"oidcs,omitempty" yaml:"oidcs,omitempty"`
}
