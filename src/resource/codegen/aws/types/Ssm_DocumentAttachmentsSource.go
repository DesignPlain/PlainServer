package types

type Ssm_DocumentAttachmentsSource struct {
	// The value of a key-value pair that identifies the location of an attachment to the document. The argument format is a list of a single string that depends on the type of key you specify - see the [API Reference](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_AttachmentsSource.html) for details.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The key of a key-value pair that identifies the location of an attachment to the document. Valid values: `SourceUrl`, `S3FileUrl`, `AttachmentReference`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The name of the document attachment file.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
