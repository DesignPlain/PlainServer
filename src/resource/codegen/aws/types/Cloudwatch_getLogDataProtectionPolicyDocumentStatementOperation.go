package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperation struct {
	// Configures the detection of sensitive data.
	Audit Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAudit `json:"audit,omitempty" yaml:"audit,omitempty"`

	/*
	   Configures the masking of sensitive data.

	   > Every policy statement must specify exactly one operation.
	*/
	Deidentify Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationDeidentify `json:"deidentify,omitempty" yaml:"deidentify,omitempty"`
}
