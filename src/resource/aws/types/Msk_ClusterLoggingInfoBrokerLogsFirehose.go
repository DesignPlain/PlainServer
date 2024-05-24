package types

type Msk_ClusterLoggingInfoBrokerLogsFirehose struct {
	// Name of the Kinesis Data Firehose delivery stream to deliver logs to.
	DeliveryStream string `json:"deliveryStream,omitempty" yaml:"deliveryStream,omitempty"`

	// Controls whether provisioned throughput is enabled or not. Default value: `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
