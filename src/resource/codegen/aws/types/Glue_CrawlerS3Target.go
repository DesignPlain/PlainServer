package types

type Glue_CrawlerS3Target struct {
	// A list of glob patterns used to exclude from the crawl.
	Exclusions []string `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`

	// The path to the Amazon S3 target.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	SampleSize int `json:"sampleSize,omitempty" yaml:"sampleSize,omitempty"`

	// The name of a connection which allows crawler to access data in S3 within a VPC.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// The ARN of the dead-letter SQS queue.
	DlqEventQueueArn string `json:"dlqEventQueueArn,omitempty" yaml:"dlqEventQueueArn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	EventQueueArn string `json:"eventQueueArn,omitempty" yaml:"eventQueueArn,omitempty"`
}
