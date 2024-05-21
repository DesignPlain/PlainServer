package types

type Container_ClusterServiceExternalIpsConfig struct {
	// Controls whether external ips specified by a service will be allowed. It is enabled by default.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
