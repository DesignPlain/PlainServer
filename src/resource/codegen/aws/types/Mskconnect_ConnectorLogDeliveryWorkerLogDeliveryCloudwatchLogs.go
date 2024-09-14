package types

type Mskconnect_ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs struct {
	// Whether log delivery to Amazon CloudWatch Logs is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the CloudWatch log group that is the destination for log delivery.
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`
}
