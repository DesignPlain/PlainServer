package types

type Oam_getLinkLinkConfiguration struct {
	// Configuration for filtering which metric namespaces are to be shared from the source account to the monitoring account. See `metric_configuration` Block for details.
	MetricConfigurations []Oam_getLinkLinkConfigurationMetricConfiguration `json:"metricConfigurations,omitempty" yaml:"metricConfigurations,omitempty"`

	// Configuration for filtering which log groups are to send log events from the source account to the monitoring account. See `log_group_configuration` Block for details.
	LogGroupConfigurations []Oam_getLinkLinkConfigurationLogGroupConfiguration `json:"logGroupConfigurations,omitempty" yaml:"logGroupConfigurations,omitempty"`
}
