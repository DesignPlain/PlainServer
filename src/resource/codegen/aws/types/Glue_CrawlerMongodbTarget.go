package types

type Glue_CrawlerMongodbTarget struct {
	// The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// The path of the Amazon DocumentDB or MongoDB target (database/collection).
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table. Default value is `true`.
	ScanAll bool `json:"scanAll,omitempty" yaml:"scanAll,omitempty"`
}
