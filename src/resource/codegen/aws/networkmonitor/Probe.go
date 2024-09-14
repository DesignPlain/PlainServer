package networkmonitor

type Probe struct {
	// The ARN of the subnet.
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`

	// Key-value tags for the monitor. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The destination IP address. This must be either IPV4 or IPV6.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The port associated with the destination. This is required only if the protocol is TCP and must be a number between 1 and 65536.
	DestinationPort int `json:"destinationPort,omitempty" yaml:"destinationPort,omitempty"`

	// The name of the monitor.
	MonitorName string `json:"monitorName,omitempty" yaml:"monitorName,omitempty"`

	/*
	   The size of the packets sent between the source and destination. This must be a number between 56 and 8500.

	   The following arguments are optional:
	*/
	PacketSize int `json:"packetSize,omitempty" yaml:"packetSize,omitempty"`

	// The protocol used for the network traffic between the source and destination. This must be either TCP or ICMP.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
