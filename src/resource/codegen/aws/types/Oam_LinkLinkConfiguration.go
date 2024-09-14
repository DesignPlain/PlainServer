package types

type Oam_LinkLinkConfiguration struct {
	// Configuration for filtering which log groups are to send log events from the source account to the monitoring account. See `log_group_configuration` Block for details.
	LogGroupConfiguration Oam_LinkLinkConfigurationLogGroupConfiguration `json:"logGroupConfiguration,omitempty" yaml:"logGroupConfiguration,omitempty"`

	// Configuration for filtering which metric namespaces are to be shared from the source account to the monitoring account. See `metric_configuration` Block for details.
	MetricConfiguration Oam_LinkLinkConfigurationMetricConfiguration `json:"metricConfiguration,omitempty" yaml:"metricConfiguration,omitempty"`
}
