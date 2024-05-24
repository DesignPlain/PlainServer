package eks

import types "DesignSphere_Server/src/resource/aws/types"

type IdentityProviderConfig struct {
	// Name of the EKS Cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster. Detailed below.
	Oidc types.Eks_IdentityProviderConfigOidc `json:"oidc,omitempty" yaml:"oidc,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
