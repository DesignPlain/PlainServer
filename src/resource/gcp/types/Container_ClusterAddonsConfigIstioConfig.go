package types

type Container_ClusterAddonsConfigIstioConfig struct {
	// The authentication type between services in Istio. Available options include `AUTH_MUTUAL_TLS`.
	Auth string `json:"auth,omitempty" yaml:"auth,omitempty"`

	/*
	   The status of the Istio addon, which makes it easy to set up Istio for services in a
	   cluster. It is disabled by default. Set `disabled = false` to enable.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}
