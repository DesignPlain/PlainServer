package types

type Ssm_DocumentAttachmentsSource struct {
	// The key describing the location of an attachment to a document. Valid key types include: `SourceUrl` and `S3FileUrl`
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The name of the document attachment file
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value describing the location of an attachment to a document
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
