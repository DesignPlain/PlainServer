package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3 struct {
	// Name of the S3 Bucket to send findings to.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
