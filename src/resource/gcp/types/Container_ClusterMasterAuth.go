package types

type Container_ClusterMasterAuth struct {
	// Base64 encoded public certificate used by clients to authenticate to the cluster endpoint.
	ClientCertificate string `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`

	// Whether client certificate authorization is enabled for this cluster.  For example:
	ClientCertificateConfig Container_ClusterMasterAuthClientCertificateConfig `json:"clientCertificateConfig,omitempty" yaml:"clientCertificateConfig,omitempty"`

	// Base64 encoded private key used by clients to authenticate to the cluster endpoint.
	ClientKey string `json:"clientKey,omitempty" yaml:"clientKey,omitempty"`

	// Base64 encoded public certificate that is the root of trust for the cluster.
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty" yaml:"clusterCaCertificate,omitempty"`
}
