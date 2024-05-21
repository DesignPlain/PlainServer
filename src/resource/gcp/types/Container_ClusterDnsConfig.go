package types

type Container_ClusterDnsConfig struct {
	// Which in-cluster DNS provider should be used. `PROVIDER_UNSPECIFIED` (default) or `PLATFORM_DEFAULT` or `CLOUD_DNS`.
	ClusterDns string `json:"clusterDns,omitempty" yaml:"clusterDns,omitempty"`

	// The suffix used for all cluster service records.
	ClusterDnsDomain string `json:"clusterDnsDomain,omitempty" yaml:"clusterDnsDomain,omitempty"`

	// The scope of access to cluster DNS records. `DNS_SCOPE_UNSPECIFIED` (default) or `CLUSTER_SCOPE` or `VPC_SCOPE`.
	ClusterDnsScope string `json:"clusterDnsScope,omitempty" yaml:"clusterDnsScope,omitempty"`
}
