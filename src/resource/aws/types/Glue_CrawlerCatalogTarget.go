package types

type Glue_CrawlerCatalogTarget struct {
	/*
	   A valid Amazon SQS ARN.

	   > --Note:-- `deletion_behavior` of catalog target doesn't support `DEPRECATE_IN_DATABASE`.

	   > --Note:-- `configuration` for catalog target crawlers will have `{ ... "Grouping": { "TableGroupingPolicy": "CombineCompatibleSchemas"} }` by default.
	*/
	DlqEventQueueArn string `json:"dlqEventQueueArn,omitempty" yaml:"dlqEventQueueArn,omitempty"`

	// A valid Amazon SQS ARN.
	EventQueueArn string `json:"eventQueueArn,omitempty" yaml:"eventQueueArn,omitempty"`

	// A list of catalog tables to be synchronized.
	Tables []string `json:"tables,omitempty" yaml:"tables,omitempty"`

	// The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a `NETWORK` Connection type.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// The name of the Glue database to be synchronized.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
