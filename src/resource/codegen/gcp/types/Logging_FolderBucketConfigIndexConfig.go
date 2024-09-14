package types

type Logging_FolderBucketConfigIndexConfig struct {
	/*
	   The LogEntry field path to index.
	   Note that some paths are automatically indexed, and other paths are not eligible for indexing. See indexing documentation for details.
	*/
	FieldPath string `json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`

	// The type of data in this index. Allowed types include `INDEX_TYPE_UNSPECIFIED`, `INDEX_TYPE_STRING` and `INDEX_TYPE_INTEGER`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
