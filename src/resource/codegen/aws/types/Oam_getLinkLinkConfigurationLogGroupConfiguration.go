package types

type Oam_getLinkLinkConfigurationLogGroupConfiguration struct {
	// Filter string that specifies  which metrics are to be shared with the monitoring account. See [MetricConfiguration](https://docs.aws.amazon.com/OAM/latest/APIReference/API_MetricConfiguration.html) for details.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
}
