package types

type Mskconnect_ConnectorLogDeliveryWorkerLogDelivery struct {
	// Details about delivering logs to Amazon CloudWatch Logs. See `cloudwatch_logs` Block for details.
	CloudwatchLogs Mskconnect_ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// Details about delivering logs to Amazon Kinesis Data Firehose. See `firehose` Block for details.
	Firehose Mskconnect_ConnectorLogDeliveryWorkerLogDeliveryFirehose `json:"firehose,omitempty" yaml:"firehose,omitempty"`

	// Details about delivering logs to Amazon S3. See `s3` Block for deetails.
	S3 Mskconnect_ConnectorLogDeliveryWorkerLogDeliveryS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
