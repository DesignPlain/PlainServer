package types

type Appmesh_VirtualNodeSpecServiceDiscoveryDns struct {
	// DNS host name for your virtual node.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// The preferred IP version that this virtual node uses. Valid values: `IPv6_PREFERRED`, `IPv4_PREFERRED`, `IPv4_ONLY`, `IPv6_ONLY`.
	IpPreference string `json:"ipPreference,omitempty" yaml:"ipPreference,omitempty"`

	// The DNS response type for the virtual node. Valid values: `LOADBALANCER`, `ENDPOINTS`.
	ResponseType string `json:"responseType,omitempty" yaml:"responseType,omitempty"`
}
