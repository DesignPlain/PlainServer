package types

type Glue_CrawlerS3Target struct {
	// The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a `NETWORK` Connection type.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	/*
	   A valid Amazon SQS ARN.

	   > --Note:-- `deletion_behavior` of catalog target doesn't support `DEPRECATE_IN_DATABASE`.

	   > --Note:-- `configuration` for catalog target crawlers will have `{ ... "Grouping": { "TableGroupingPolicy": "CombineCompatibleSchemas"} }` by default.
	*/
	DlqEventQueueArn string `json:"dlqEventQueueArn,omitempty" yaml:"dlqEventQueueArn,omitempty"`

	// A valid Amazon SQS ARN.
	EventQueueArn string `json:"eventQueueArn,omitempty" yaml:"eventQueueArn,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []string `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	SampleSize int `json:"sampleSize,omitempty" yaml:"sampleSize,omitempty"`
}
