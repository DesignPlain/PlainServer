package types

type Edgecontainer_ClusterSystemAddonsConfigIngress struct {
	// Whether Ingress is disabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Ingress VIP.
	Ipv4Vip string `json:"ipv4Vip,omitempty" yaml:"ipv4Vip,omitempty"`
}
