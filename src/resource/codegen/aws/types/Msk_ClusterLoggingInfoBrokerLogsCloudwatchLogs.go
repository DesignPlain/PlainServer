package types

type Msk_ClusterLoggingInfoBrokerLogsCloudwatchLogs struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Name of the Cloudwatch Log Group to deliver logs to.
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`
}
