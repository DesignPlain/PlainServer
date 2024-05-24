package types

type Msk_ClusterLoggingInfoBrokerLogs struct {
	//
	CloudwatchLogs Msk_ClusterLoggingInfoBrokerLogsCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	//
	Firehose Msk_ClusterLoggingInfoBrokerLogsFirehose `json:"firehose,omitempty" yaml:"firehose,omitempty"`

	//
	S3 Msk_ClusterLoggingInfoBrokerLogsS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
