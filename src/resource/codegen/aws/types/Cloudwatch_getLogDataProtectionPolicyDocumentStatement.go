package types

type Cloudwatch_getLogDataProtectionPolicyDocumentStatement struct {
	// Name of this statement.
	Sid string `json:"sid,omitempty" yaml:"sid,omitempty"`

	// Set of at least 1 sensitive data identifiers that you want to mask. Read more in [Types of data that you can protect](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/protect-sensitive-log-data-types.html).
	DataIdentifiers []string `json:"dataIdentifiers,omitempty" yaml:"dataIdentifiers,omitempty"`

	// Configures the data protection operation applied by this statement.
	Operation Cloudwatch_getLogDataProtectionPolicyDocumentStatementOperation `json:"operation,omitempty" yaml:"operation,omitempty"`
}
