package types

type Msk_ClusterLoggingInfoBrokerLogsFirehose struct {
	// Name of the Kinesis Data Firehose delivery stream to deliver logs to.
	DeliveryStream string `json:"deliveryStream,omitempty" yaml:"deliveryStream,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
