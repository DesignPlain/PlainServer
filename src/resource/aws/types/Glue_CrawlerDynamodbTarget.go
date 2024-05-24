package types

type Glue_CrawlerDynamodbTarget struct {
	// The name of the DynamoDB table to crawl.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to `true`.
	ScanAll bool `json:"scanAll,omitempty" yaml:"scanAll,omitempty"`

	// The percentage of the configured read capacity units to use by the AWS Glue crawler. The valid values are null or a value between 0.1 to 1.5.
	ScanRate float64 `json:"scanRate,omitempty" yaml:"scanRate,omitempty"`
}
