package types

type Kendra_IndexIndexStatisticTextDocumentStatistic struct {
	// The total size, in bytes, of the indexed documents.
	IndexedTextBytes int `json:"indexedTextBytes,omitempty" yaml:"indexedTextBytes,omitempty"`

	// The number of text documents indexed.
	IndexedTextDocumentsCount int `json:"indexedTextDocumentsCount,omitempty" yaml:"indexedTextDocumentsCount,omitempty"`
}
