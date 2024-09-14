package types

type Securitylake_CustomLogSourceAttribute struct {
	// The ARN of the AWS Glue crawler.
	CrawlerArn string `json:"crawlerArn,omitempty" yaml:"crawlerArn,omitempty"`

	// The ARN of the AWS Glue database where results are written.
	DatabaseArn string `json:"databaseArn,omitempty" yaml:"databaseArn,omitempty"`

	// The ARN of the AWS Glue table.
	TableArn string `json:"tableArn,omitempty" yaml:"tableArn,omitempty"`
}
