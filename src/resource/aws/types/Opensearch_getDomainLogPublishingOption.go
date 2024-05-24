package types

type Opensearch_getDomainLogPublishingOption struct {
	// CloudWatch Log Group where the logs are published.
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn,omitempty" yaml:"cloudwatchLogGroupArn,omitempty"`

	// Enabled disabled toggle for off-peak update window
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Type of OpenSearch log being published.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`
}
