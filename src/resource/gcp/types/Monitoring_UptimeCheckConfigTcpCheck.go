package types

type Monitoring_UptimeCheckConfigTcpCheck struct {
	/*
	   Contains information needed to add pings to a TCP check.
	   Structure is documented below.
	*/
	PingConfig Monitoring_UptimeCheckConfigTcpCheckPingConfig `json:"pingConfig,omitempty" yaml:"pingConfig,omitempty"`

	// The port to the page to run the check against. Will be combined with host (specified within the `monitored_resource`) to construct the full URL.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
