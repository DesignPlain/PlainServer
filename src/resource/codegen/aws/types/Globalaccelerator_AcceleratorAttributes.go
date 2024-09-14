package types

type Globalaccelerator_AcceleratorAttributes struct {
	// The prefix for the location in the Amazon S3 bucket for the flow logs. Required if `flow_logs_enabled` is `true`.
	FlowLogsS3Prefix string `json:"flowLogsS3Prefix,omitempty" yaml:"flowLogsS3Prefix,omitempty"`

	// Indicates whether flow logs are enabled. Defaults to `false`. Valid values: `true`, `false`.
	FlowLogsEnabled bool `json:"flowLogsEnabled,omitempty" yaml:"flowLogsEnabled,omitempty"`

	// The name of the Amazon S3 bucket for the flow logs. Required if `flow_logs_enabled` is `true`.
	FlowLogsS3Bucket string `json:"flowLogsS3Bucket,omitempty" yaml:"flowLogsS3Bucket,omitempty"`
}
