package types

type Kinesis_FirehoseDeliveryStreamOpensearchConfigurationDocumentIdOptions struct {
	// The method for setting up document ID. Valid values: `FIREHOSE_DEFAULT`, `NO_DOCUMENT_ID`.
	DefaultDocumentIdFormat string `json:"defaultDocumentIdFormat,omitempty" yaml:"defaultDocumentIdFormat,omitempty"`
}
