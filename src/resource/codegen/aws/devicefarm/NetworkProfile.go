package devicefarm

type NetworkProfile struct {
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	DownlinkJitterMs int `json:"downlinkJitterMs,omitempty" yaml:"downlinkJitterMs,omitempty"`

	// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	UplinkDelayMs int `json:"uplinkDelayMs,omitempty" yaml:"uplinkDelayMs,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	DownlinkDelayMs int `json:"downlinkDelayMs,omitempty" yaml:"downlinkDelayMs,omitempty"`

	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	DownlinkBandwidthBits int `json:"downlinkBandwidthBits,omitempty" yaml:"downlinkBandwidthBits,omitempty"`

	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	DownlinkLossPercent int `json:"downlinkLossPercent,omitempty" yaml:"downlinkLossPercent,omitempty"`

	// The name for the network profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of the project for the network profile.
	ProjectArn string `json:"projectArn,omitempty" yaml:"projectArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	UplinkBandwidthBits int `json:"uplinkBandwidthBits,omitempty" yaml:"uplinkBandwidthBits,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	UplinkJitterMs int `json:"uplinkJitterMs,omitempty" yaml:"uplinkJitterMs,omitempty"`

	// The description of the network profile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	UplinkLossPercent int `json:"uplinkLossPercent,omitempty" yaml:"uplinkLossPercent,omitempty"`
}
