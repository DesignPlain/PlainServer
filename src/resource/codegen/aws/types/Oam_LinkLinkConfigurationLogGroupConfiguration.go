package types

type Oam_LinkLinkConfigurationLogGroupConfiguration struct {
	// Filter string that specifies which log groups are to share their log events with the monitoring account. See [LogGroupConfiguration](https://docs.aws.amazon.com/OAM/latest/APIReference/API_LogGroupConfiguration.html) for details.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
}
