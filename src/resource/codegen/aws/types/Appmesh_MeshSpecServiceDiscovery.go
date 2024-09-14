package types

type Appmesh_MeshSpecServiceDiscovery struct {
	// The IP version to use to control traffic within the mesh. Valid values are `IPv6_PREFERRED`, `IPv4_PREFERRED`, `IPv4_ONLY`, and `IPv6_ONLY`.
	IpPreference string `json:"ipPreference,omitempty" yaml:"ipPreference,omitempty"`
}
