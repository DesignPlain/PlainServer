package types

type Elasticache_ClusterLogDeliveryConfiguration struct {
	// Name of either the CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// For CloudWatch Logs use `cloudwatch-logs` or for Kinesis Data Firehose use `kinesis-firehose`.
	DestinationType string `json:"destinationType,omitempty" yaml:"destinationType,omitempty"`

	// Valid values are `json` or `text`
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`

	// Valid values are  `slow-log` or `engine-log`. Max 1 of each.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`
}
