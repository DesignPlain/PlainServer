package types

type Lambda_EventSourceMappingDocumentDbEventSourceConfig struct {
	// The name of the collection to consume within the database. If you do not specify a collection, Lambda consumes all collections.
	CollectionName string `json:"collectionName,omitempty" yaml:"collectionName,omitempty"`

	// The name of the database to consume within the DocumentDB cluster.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Determines what DocumentDB sends to your event stream during document update operations. If set to `UpdateLookup`, DocumentDB sends a delta describing the changes, along with a copy of the entire document. Otherwise, DocumentDB sends only a partial document that contains the changes. Valid values: `UpdateLookup`, `Default`.
	FullDocument string `json:"fullDocument,omitempty" yaml:"fullDocument,omitempty"`
}
