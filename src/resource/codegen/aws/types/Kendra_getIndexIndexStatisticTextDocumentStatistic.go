package types

type Kendra_getIndexIndexStatisticTextDocumentStatistic struct {
	// The number of text documents indexed.
	IndexedTextDocumentsCount int `json:"indexedTextDocumentsCount,omitempty" yaml:"indexedTextDocumentsCount,omitempty"`

	// Total size, in bytes, of the indexed documents.
	IndexedTextBytes int `json:"indexedTextBytes,omitempty" yaml:"indexedTextBytes,omitempty"`
}
