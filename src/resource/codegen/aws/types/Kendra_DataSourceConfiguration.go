package types

type Kendra_DataSourceConfiguration struct {
	// A block that provides the configuration information to connect to an Amazon S3 bucket as your data source. Detailed below.
	S3Configuration Kendra_DataSourceConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// A block that provides the configuration information required for Amazon Kendra Web Crawler. Detailed below.
	WebCrawlerConfiguration Kendra_DataSourceConfigurationWebCrawlerConfiguration `json:"webCrawlerConfiguration,omitempty" yaml:"webCrawlerConfiguration,omitempty"`
}
