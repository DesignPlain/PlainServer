package types

type Opensearch_DomainLogPublishingOption struct {
	// ARN of the Cloudwatch log group to which log needs to be published.
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn,omitempty" yaml:"cloudwatchLogGroupArn,omitempty"`

	// Whether given log publishing option is enabled or not.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Type of OpenSearch log. Valid values: `INDEX_SLOW_LOGS`, `SEARCH_SLOW_LOGS`, `ES_APPLICATION_LOGS`, `AUDIT_LOGS`.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`
}
