package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs struct {
	// Name of the CloudWatch Log Group to send findings to.
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`
}
