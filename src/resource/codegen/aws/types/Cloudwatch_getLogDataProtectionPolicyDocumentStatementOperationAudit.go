package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAudit struct {
	// Configures destinations to send audit findings to.
	FindingsDestination Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination `json:"findingsDestination,omitempty" yaml:"findingsDestination,omitempty"`
}
