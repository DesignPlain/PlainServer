package types

type Mskconnect_ConnectorLogDeliveryWorkerLogDeliveryFirehose struct {
	// The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
	DeliveryStream string `json:"deliveryStream,omitempty" yaml:"deliveryStream,omitempty"`

	// Specifies whether connector logs get delivered to Amazon Kinesis Data Firehose.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
