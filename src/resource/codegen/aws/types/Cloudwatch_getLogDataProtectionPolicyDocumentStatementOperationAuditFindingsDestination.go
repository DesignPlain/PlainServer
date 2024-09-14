package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination struct {
	// Configures CloudWatch Logs as a findings destination.
	CloudwatchLogs Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// Configures Kinesis Firehose as a findings destination.
	Firehose Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose `json:"firehose,omitempty" yaml:"firehose,omitempty"`

	// Configures S3 as a findings destination.
	S3 Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
