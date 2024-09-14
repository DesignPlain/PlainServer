package networkmonitor

type Monitor struct {
	/*
	   The name of the monitor.

	   The following arguments are optional:
	*/
	MonitorName string `json:"monitorName,omitempty" yaml:"monitorName,omitempty"`

	// Key-value tags for the monitor. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The time, in seconds, that metrics are aggregated and sent to Amazon CloudWatch. Valid values are either 30 or 60.
	AggregationPeriod int `json:"aggregationPeriod,omitempty" yaml:"aggregationPeriod,omitempty"`
}
