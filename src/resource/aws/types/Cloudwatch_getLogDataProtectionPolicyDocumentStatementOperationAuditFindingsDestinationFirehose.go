package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose struct {
	// Name of the Kinesis Firehose Delivery Stream to send findings to.
	DeliveryStream string `json:"deliveryStream,omitempty" yaml:"deliveryStream,omitempty"`
}
