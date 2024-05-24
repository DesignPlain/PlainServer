package types

type Elasticsearch_getDomainLogPublishingOption struct {
	// The CloudWatch Log Group where the logs are published.
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn,omitempty" yaml:"cloudwatchLogGroupArn,omitempty"`

	// Whether node to node encryption is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The type of Elasticsearch log being published.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`
}
