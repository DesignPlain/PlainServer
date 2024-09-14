package types

type Container_getClusterDnsConfig struct {
	// The suffix used for all cluster service records.
	ClusterDnsDomain string `json:"clusterDnsDomain,omitempty" yaml:"clusterDnsDomain,omitempty"`

	// The scope of access to cluster DNS records.
	ClusterDnsScope string `json:"clusterDnsScope,omitempty" yaml:"clusterDnsScope,omitempty"`

	// Which in-cluster DNS provider should be used.
	ClusterDns string `json:"clusterDns,omitempty" yaml:"clusterDns,omitempty"`
}
