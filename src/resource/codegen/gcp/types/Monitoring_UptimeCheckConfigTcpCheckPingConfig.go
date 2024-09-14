package types

type Monitoring_UptimeCheckConfigTcpCheckPingConfig struct {
	// Number of ICMP pings. A maximum of 3 ICMP pings is currently supported.
	PingsCount int `json:"pingsCount,omitempty" yaml:"pingsCount,omitempty"`
}
